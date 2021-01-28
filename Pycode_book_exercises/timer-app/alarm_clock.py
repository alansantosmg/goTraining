import time

current_time = time.localtime()
hour = current_time.tm_hour
#minute = current_time.tm_minute
if (hour > 7) or (hour==7 and minute > 29): 
    print('Hora de acordar')
  