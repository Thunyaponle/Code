void setup() {
  pinMode(BCM0,OUTPUT);
  pinMode(BCM1,OUTPUT);
}

void loop() {
  digitalWrite(BCM0,1);
  digitalWrite(BCM1,1);
  delay(500);
  digitalWrite(BCM0,0);
  digitalWrite(BCM1,0);
  delay(500);
}
