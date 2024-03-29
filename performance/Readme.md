Overview
========

Concurrent Access & Lock
------------------------

Locks helps to avoid concurrent access to resources.
However, it downgrades the app's performance because
of the contention issue. We can
narrow down the range of a lock to reduce this contention.

[Atomic][atomic] gives a comparion of performance
between atomic and mutex lock.

Sync.Pool
---------

In the garbage-collected world, we want to
keep the GC overhead as little as possible.
One of the things we can do is limiting
the number of allocations in our application.

***[sync.Pool][sync.Pool]***
caches allocated but unused items for later reuse,
relieving pressure on the garbage collector.
The real power of it is visible when there're
frequent allocations and deallocations of
the same data structure
(especially when the structure object is expensive to create).

The article [Using Sync.Pool][Using Sync.Pool] gives
a clear explanation of the power of the pool.

[sync.Pool]: https://pkg.go.dev/sync#Pool
[Using Sync.Pool]: https://developer20.com/using-sync-pool/
[atomic]: ./atomic
