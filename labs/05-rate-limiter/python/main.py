import threading
import time

class RateLimiter:
    def __init__(self, capacity, refill_amount, refill_interval, clock=time.monotonic):
        if capacity <= 0:
            raise ValueError("Capacity must be greater than 0")
        if refill_amount <= 0:
            raise ValueError("Refill amount must be greater than 0")
        if refill_interval <= 0:
            raise ValueError("Refill interval must be greater than 0")
        self.capacity = capacity
        self.tokens = capacity
        self.refill_amount = refill_amount
        self.refill_interval = refill_interval
        self.clock = clock
        self.last_refill = self.clock()
        self.lock = threading.Lock()

    def allow(self):
        with self.lock:
            current_time = self.clock()
            elapsed_time = current_time - self.last_refill
            if elapsed_time >= self.refill_interval:
                refill_count = int(elapsed_time / self.refill_interval)
                self.tokens = min(self.capacity, self.tokens + refill_count * self.refill_amount)
                self.last_refill += refill_count * self.refill_interval

            if self.tokens > 0:
                self.tokens -= 1
                return True
            else:
                return False
        