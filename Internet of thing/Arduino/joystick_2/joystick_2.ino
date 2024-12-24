//650710554 ธันยพร แซ่ลี้
#include "TFT_eSPI.h"
TFT_eSPI tft;
void setup() {
  pinMode(WIO_5S_LEFT, INPUT_PULLUP);

  pinMode(WIO_5S_RIGHT, INPUT_PULLUP);

  tft.begin();
  tft.setRotation(3);
  tft.fillScreen(TFT_BLACK);
  tft.setTextColor(TFT_WHITE);
  tft.setTextSize(3);
  tft.drawString("assignment 2",50,10);
}

void loop() {
  int location = 290;  //min =0

  while (true) {

    if (digitalRead(WIO_5S_LEFT) == LOW) {

      int startRowRight = location;
      tft.fillRoundRect(startRowRight, 110, 25, 100, 10, TFT_WHITE);

      delay(500);

      if (location -= 27 > 0) {
        location -= 27;
      }

    } else if (digitalRead(WIO_5S_RIGHT) == LOW) {

      if (location += 27 < 319) {
        location += 27;
      }

      int startRowLeft = location;
      tft.fillRoundRect(startRowLeft, 110, 25, 100, 10, TFT_BLACK);

      delay(500);
    }
  }
}
