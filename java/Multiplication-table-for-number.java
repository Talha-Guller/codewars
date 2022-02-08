//https://www.codewars.com/kata/5a2fd38b55519ed98f0000ce/train/java

class Kata {
    public static String multiTable(int num){
        String result = "";
        for(int i = 1; i <= 10; i++) {
            result += i+" * " + num+" = " + (i * num + "\n");
        }

       return result.substring(0, result.length() - 1);
    }
}
