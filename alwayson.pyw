import ctypes

ctypes.windll.kernel32.SetThreadExecutionState(0x80000002)
ctypes.windll.user32.MessageBoxW(0,u"exit?",u"alwayson",0)
