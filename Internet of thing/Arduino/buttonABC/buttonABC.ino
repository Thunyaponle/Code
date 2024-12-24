#include <TFT_eSPI.h>
TFT_eSPI tft;

void setup() {
  // put your setup code here, to run once:
  pinMode(WIO_KEY_A,INPUT_PULLUP);
  pinMode(WIO_KEY_B,INPUT_PULLUP);
  pinMode(WIO_KEY_C,INPUT_PULLUP);
  tft.begin();
  tft.setRotation(3);
  tft.setTextColor(TFT_WHITE,TFT_BLACK);
  tft.setTextSize(2);
  tft.drawString("Press any button!",60,80);
  tft.setTextColor(TFT_GREEN,TFT_BLACK);
}

void loop() {
  // put your main code here, to run repeatedly:
  if(digitalRead(WIO_KEY_A)==LOW){
    tft.drawString("button A Pressed",60,120);
  }
  else if(digitalRead(WIO_KEY_B)==LOW){
    tft.drawString("button B Pressed",60,120);
  }
  else if(digitalRead(WIO_KEY_C)==LOW){
    tft.drawString("button C Pressed",60,120);
  }
  }
  