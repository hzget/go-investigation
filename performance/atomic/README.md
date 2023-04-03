Atomic VS Mutex Lock
====================

Locks helps to avoid concurrent access to resources.
However, it downgrades the app's performance because
of the [contention issue][resource contention]. We can
narrow down the range of a lock to reduce this contention.

[bigcache][bigcache] gives an example. It splits the
resource into N shards with each contains its own instance
of the cache with a lock. When a goroutine access one shard,
it will not block other goroutines from accessing to other shards.

[Atomic][go pkg atomic] just adds a lock in the CPU instruction
level --- a very tiny range of the lock.

The commit history for the go pkg [context][go context ae1fa08]
shows that it ever improves
their performance via replacing mutex lock operation by atomic operation.

[main_test.go](./main_test.go) gives a comparison of performance
between atomic and mutex lock.

    > benchstat mutex.benchmark atomic.benchmark 
    name             old time/op    new time/op    delta
    IncVisitCount-4    38.1ns ± 1%    17.8ns ± 4%  -53.23%  (p=0.000 n=7+10)

    name             old alloc/op   new alloc/op   delta
    IncVisitCount-4     0.00B          0.00B          ~     (all equal)

    name             old allocs/op  new allocs/op  delta
    IncVisitCount-4      0.00           0.00          ~     (all equal)

Underlying Principle of Atomic
------------------------------

The previous example makes use of Add() func in the pkg
[atomic][go pkg atomic].
If we dig further into the [instruction code][atomic_amd64.s],
we can find that it makes use
of the `LOCK` instruction. (for amd64 platform):

    // src/sync/atomic/type.go
    // Add atomically adds delta to x and returns the new value.
    func (x *Uint64) Add(delta uint64) (new uint64) { return AddUint64(&x.v, delta) }

    // src/sync/atomic/asm.s
    TEXT ·AddUint64(SB),NOSPLIT,$0
      JMP  runtime∕internal∕atomic·Xadd64(SB)

    // in runtime/internal/atomic/atomic_amd64.s
    // uint64 Xadd64(uint64 volatile *val, int64 delta)
    // Atomically:
    //  *val += delta;
    //  return *val;
    TEXT ·Xadd64(SB), NOSPLIT, $0-24
        MOVQ    ptr+0(FP), BX
        MOVQ    delta+8(FP), AX
        MOVQ    AX, CX
        LOCK              <------- here is the `LOCK` instruction
        XADDQ   AX, 0(BX)
        ADDQ    CX, AX
        MOVQ    AX, ret+16(FP)
        RET

According to the [x86 and amd64 instruction reference][lock instruction],
we have that
> Causes the processor’s LOCK# signal to be asserted during
> execution of the accompanying instruction (turns the instruction
> into an **atomic instruction**). In a multiprocessor environment,
> the LOCK# signal ensures that the processor has exclusive use of
> any shared memory while the signal is asserted.

[go pkg atomic]: https://pkg.go.dev/sync/atomic
[atomic_amd64.s]: https://cs.opensource.google/go/go/+/refs/tags/go1.20.2:src/runtime/internal/atomic/atomic_amd64.s
[lock instruction]: https://www.felixcloutier.com/x86/lock
[resource contention]: https://en.wikipedia.org/wiki/Resource_contention
[bigcache]: https://blog.allegro.tech/2016/03/writing-fast-cache-service-in-go.html#concurrency
[go context ae1fa08]: https://cs.opensource.google/go/go/+/ae1fa08e4138c49c8e7fa10c3eadbfca0233842b
