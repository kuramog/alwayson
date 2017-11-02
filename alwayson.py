import ctypes,time

ctypes.windll.kernel32.SetThreadExecutionState(0x80000002)
while True:
    try:
        time.sleep(2)
    except KeyboardInterrupt:
        break
