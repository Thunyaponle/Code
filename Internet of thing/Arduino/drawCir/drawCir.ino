#include <TFT_eSPI.h>
TFT_eSPI tft;

void setup() {
  // put your setup code here, to run once:
  tft.begin();
  tft.setRotation(3);
  tft.fillScreen(TFT_BLUE);
  for(int i=10;i<120;i+=5){
    tft.drawCircle(180,120,i,TFT_YELLOW);
  }
  
}

void loop() {
  // put your main code here, to run repeatedly:
  
}