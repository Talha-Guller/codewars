# https://www.codewars.com/kata/55192f4ecd82ff826900089e/train/java

public class Kata {
  
  public static boolean divide(int weight) {
        if (1<=weight){
            if (weight<=100){
                if (weight % 2 == 0){
                    if (weight == 2){
                        return false;
                    }
                    return true;
                }
            }
        }
        return false;
    }
}
