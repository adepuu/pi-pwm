import RPi.GPIO as GPIO
import time
import signal
import sys
from rpi_hardware_pwm import HardwarePWM

WAIT_TIME = 1 
FAN_PIN = 12
PWM_FREQ = 25

PWM = HardwarePWM(pwm_channel=0, hz=25_000)
time.sleep(WAIT_TIME)

try:
    signal.signal(signal.SIGTERM, lambda *args: sys.exit(0))
    GPIO.setwarnings(True)
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(17, GPIO.OUT, initial=GPIO.HIGH)
    GPIO.setup(27, GPIO.OUT, initial=GPIO.HIGH)
    GPIO.setup(10, GPIO.OUT, initial=GPIO.HIGH)
    GPIO.setup(22, GPIO.OUT, initial=GPIO.HIGH)
    PWM.start(10)
    PWM.change_duty_cycle(10)

    while True:
        PWM.change_duty_cycle(10)
        # for dc in range(10, 101, 5):
        #     print(dc)
        #     PWM.change_duty_cycle(dc)
        #     time.sleep(WAIT_TIME)
        # for dc in range(100, 10, -5):
        #     print(dc)
        #     PWM.change_duty_cycle(dc)
        #     time.sleep(WAIT_TIME)

except KeyboardInterrupt:
    pass

finally:
    PWM.stop()
    GPIO.cleanup()