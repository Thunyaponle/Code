//650710554 ธันยพร แซ่ลี้
#include "TFT_eSPI.h"
TFT_eSPI tft;
void setup() {
  pinMode(WIO_5S_LEFT, INPUT_PULLUP);

  pinMode(WIO_5S_RIGHT, INPUT_PULLUP);

  pinMode(BCM1, OUTPUT);
  pinMode(BCM0, OUTPUT);

  tft.begin();
  tft.setRotation(3);
  tft.fillScreen(TFT_BLACK);
  tft.setTextColor(TFT_WHITE);
  tft.setTextSize(3);
  tft.drawString("assignment 3", 50, 10);
}

void loop() {
  //int time = 1000;
  int end = 319;
  int location = 0;
  while (true) {
    int startRowRight = location;
    

    digitalWrite(BCM0, 1);
    digitalWrite(BCM1, 0);
    delay(end - startRowRight);

    digitalWrite(BCM0, 0);
    digitalWrite(BCM1, 1);
    delay(end - startRowRight);

    if (digitalRead(WIO_5S_RIGHT) == LOW) {
      tft.fillRoundRect(startRowRight, 110, 25, 100, 10, TFT_WHITE);



      if (location += 27 < 319) {
        location += 27;
      }

    } else if (digitalRead(WIO_5S_LEFT) == LOW) {
      if (location -= 27 > 0) {
        location -= 27;
      }
      int startRowLeft = location;
      tft.fillRoundRect(startRowLeft, 110, 25, 100, 10, TFT_BLACK);

      
    }
  }
}
