#include <TFT_eSPI.h>
TFT_eSPI tft;

void setup() {
  pinMode(A0,INPUT);
  tft.begin();
  tft.setRotation(3);
  tft.setTextSize(3);
  tft.fillScreen(TFT_BLACK);
  tft.setTextColor(TFT_WHITE,TFT_BLACK);
  tft.drawString("ADC GPIO27",60,60);
}

void loop() {
  int adc = analogRead(A0);
  tft.drawString("ADC0:"+String(adc)+String(" "),60,120);
  delay(100);

}
