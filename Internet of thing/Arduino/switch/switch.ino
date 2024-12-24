void setup() {
  pinMode(BCM0, INPUT);
  pinMode(BCM1, OUTPUT);
}

void loop() {
  if (!digitalRead(BCM0))  //ถ้าไม่ใช่ 1 ให้ทำตามใน if
  {
    digitalWrite(BCM1,1);

  } else {
    digitalWrite(BCM1,0);
    delay(200);
  }
}