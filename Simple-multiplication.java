# https://www.codewars.com/kata/583710ccaa6717322c000105/train/java


public class Sid {
     public static int simpleMultiplication(int n) {
        int sonuc = 0;
        if ((n % 2)==0){
            sonuc = n * 8;
        }else{
           sonuc = n * 9;
        }
        return sonuc;
    }
}
