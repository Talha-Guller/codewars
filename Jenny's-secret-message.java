# https://www.codewars.com/kata/55225023e1be1ec8bc000390/train/java

public class Greeter {
  public static String greet(String name) {
        if(name == "Johnny"){
            return "Hello, my love!";
        }else
        return String.format("Hello, %s!", name);
    }
}
