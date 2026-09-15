export function capacityPercent(depth, capacity) {
  if (!Number.isFinite(depth) || !Number.isFinite(capacity) || capacity <= 0) return 0;
  return Math.min(100, Math.max(0, Math.round((depth / capacity) * 100)));
}

export function formatDuration(milliseconds) {
  const seconds = Math.max(0, Math.floor(milliseconds / 1000));
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const remainder = seconds % 60;
  return [hours, minutes, remainder].map((value) => String(value).padStart(2, "0")).join(":");
}

export class HttpDashboardClient {
  constructor(baseURL = "", fetcher = fetch) {
    this.baseURL = baseURL;
    this.fetcher = fetcher;
  }

  async getSnapshot() {
    return this.#request("/api/snapshot");
  }

  async queueEvent(event) {
    return this.#request("/api/events", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(event),
    });
  }

  async #request(path, init) {
    const response = await this.fetcher(`${this.baseURL}${path}`, init);
    const body = await response.json();
    if (!response.ok) throw new Error(body.message || `Request failed (${response.status})`);
    return body;
  }
}

export class DashboardView {
  constructor(root = document) {
    this.root = root;
  }

  render(snapshot) {
    const percent = capacityPercent(snapshot.queue.depth, snapshot.queue.capacity);
    this.#text("status-label", snapshot.health.status);
    this.#node("status").dataset.status = snapshot.health.status;
    this.#text("queue-ratio", `${percent}%`);
    this.#text("queue-depth", snapshot.queue.depth);
    this.#text("queue-capacity", snapshot.queue.capacity);
    this.#node("queue-meter").style.width = `${percent}%`;
    this.#text("worker-active", snapshot.workers.active);
    this.#text("worker-processed", snapshot.workers.processed);
    this.#text("worker-failed", snapshot.workers.failed);
    this.#text("uptime", formatDuration(snapshot.uptimeMs));
    this.#text("captured-at", new Date(snapshot.capturedAt).toLocaleTimeString());
  }

  setConnectionError(error) {
    this.#text("status-label", "Offline");
    this.#node("status").dataset.status = "offline";
    this.showNotice(error.message, true);
  }

  showNotice(message, isError = false) {
    const notice = this.#node("notice");
    notice.textContent = message;
    notice.dataset.error = String(isError);
  }

  eventForm() {
    return this.#node("event-form");
  }

  #node(role) {
    return this.root.querySelector(`[data-role="${role}"]`);
  }

  #text(role, value) {
    this.#node(role).textContent = String(value);
  }
}

export class DashboardController {
  constructor(client, view, pollInterval = 1500) {
    this.client = client;
    this.view = view;
    this.pollInterval = pollInterval;
    this.timer = undefined;
  }

  start() {
    this.view.eventForm().addEventListener("submit", (event) => this.submit(event));
    this.refresh();
    this.timer = setInterval(() => this.refresh(), this.pollInterval);
  }

  stop() {
    clearInterval(this.timer);
  }

  async refresh() {
    try {
      this.view.render(await this.client.getSnapshot());
    } catch (error) {
      this.view.setConnectionError(error);
    }
  }

  async submit(event) {
    event.preventDefault();
    const form = event.currentTarget;
    const button = form.querySelector("button");
    const fields = new FormData(form);
    button.disabled = true;
    try {
      const result = await this.client.queueEvent({
        source: fields.get("source"),
        payload: fields.get("payload"),
      });
      this.view.showNotice(`${result.id} accepted by the queue`);
      await this.refresh();
    } catch (error) {
      this.view.showNotice(error.message, true);
    } finally {
      button.disabled = false;
    }
  }
}

export function startDashboard(root = document) {
  const controller = new DashboardController(new HttpDashboardClient(), new DashboardView(root));
  controller.start();
  return controller;
}

if (typeof document !== "undefined") startDashboard();
