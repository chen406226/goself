import com.sun.deploy.util.StringUtils;

import java.util.*;

public class Test {
    static String table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF";
    static  int[] s = new int[]{11, 10, 3, 8, 4, 6};
    static long xor = 177451812L;
    static long add = 8728348608L;
    static private Map<Character, Integer> tr = new HashMap<>();
    public static void main(String[] args) {
        for (int i = 0; i < 58; i++) {
            tr.put(table.charAt(i), i);
        }
        System.out.println(dec("BV17x411w7KC"));
        System.out.println(dec("BV1Q541167Qg"));
        System.out.println(dec("BV1mK4y1C7Bz"));
        System.out.println(enc(170001));
        System.out.println(enc(455017605));
        System.out.println(enc(882584971));
//        System.out.println("dec=" + dec("BV17x411w7KC"));
//        System.out.println("enc=" + enc(170001L));
    }
    static public long dec(String x) {
        long r = 0;
        for (int i = 0; i < 6; i++) {
            r += tr.get(x.charAt(s[i])) * Math.pow(58, i);
        }
        return (r - add) ^ xor;
    }
    static public String enc(long xa) {
        long x = (xa ^ xor) + add;
        String[] r = "BV1  4 1 7  ".split("");
        for (int i = 0; i < 6; i++) {
            r[s[i]]= String.valueOf(table.charAt((int) ((x/(Math.pow(58, i)))%58)));
        }
        return StringUtils.join(Arrays.asList(r),"");
    }
}
