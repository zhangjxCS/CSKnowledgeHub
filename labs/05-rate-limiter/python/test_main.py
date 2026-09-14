import threading
import unittest
from concurrent.futures import ThreadPoolExecutor

from main import RateLimiter


class FakeClock:
    def __init__(self):
        self.now = 0

    def __call__(self):
        return self.now

    def advance(self, seconds):
        self.now += seconds


class TestRateLimiter(unittest.TestCase):
    def setUp(self):
        self.clock = FakeClock()

    def test_initial_capacity_and_exhaustion(self):
        limiter = RateLimiter(3, 1, 10, self.clock)

        self.assertEqual([limiter.allow() for _ in range(4)], [True, True, True, False])

    def test_refill_at_exact_boundary(self):
        limiter = RateLimiter(1, 1, 10, self.clock)
        self.assertTrue(limiter.allow())

        self.clock.advance(9)
        self.assertFalse(limiter.allow())

        self.clock.advance(1)
        self.assertTrue(limiter.allow())

    def test_invalid_configuration(self):
        for arguments in [(0, 1, 1), (1, 0, 1), (1, 1, 0)]:
            with self.subTest(arguments=arguments):
                with self.assertRaises(ValueError):
                    RateLimiter(*arguments, clock=self.clock)

    def test_concurrent_callers_do_not_share_tokens(self):
        capacity = 5
        caller_count = 20
        limiter = RateLimiter(capacity, 1, 10, self.clock)
        barrier = threading.Barrier(caller_count)

        def call_allow(_):
            barrier.wait()
            return limiter.allow()

        with ThreadPoolExecutor(max_workers=caller_count) as executor:
            results = list(executor.map(call_allow, range(caller_count)))

        self.assertEqual(sum(results), capacity)


if __name__ == "__main__":
    unittest.main()
