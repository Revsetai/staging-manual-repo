import test from "node:test";
import assert from "node:assert/strict";

import { HttpDashboardClient, capacityPercent, formatDuration } from "./app.mjs";

test("capacityPercent clamps queue pressure", () => {
  assert.equal(capacityPercent(4, 10), 40);
  assert.equal(capacityPercent(12, 10), 100);
  assert.equal(capacityPercent(1, 0), 0);
});

test("formatDuration renders a stable operations clock", () => {
  assert.equal(formatDuration(3_661_000), "01:01:01");
});

test("HttpDashboardClient posts an event", async () => {
  const calls = [];
  const client = new HttpDashboardClient("https://ingest.test", async (...args) => {
    calls.push(args);
    return { ok: true, json: async () => ({ id: "evt-1", status: "queued" }) };
  });

  const result = await client.queueEvent({ source: "test", payload: "{}" });

  assert.equal(result.id, "evt-1");
  assert.equal(calls[0][0], "https://ingest.test/api/events");
  assert.equal(calls[0][1].method, "POST");
});

test("HttpDashboardClient exposes API errors", async () => {
  const client = new HttpDashboardClient("", async () => ({
    ok: false,
    status: 429,
    json: async () => ({ message: "the queue is at capacity" }),
  }));

  await assert.rejects(() => client.queueEvent({}), /queue is at capacity/);
});
