import java.io.File;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Scanner;
import java.util.stream.Collectors;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

class Solution01 {
    public static void main(String[] args) {
        try {
            new Solution01().run();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }    

    private void run() throws FileNotFoundException {
        File f = new File("input");
        Scanner sc = new Scanner(f);

        List<String> lines = new ArrayList<>();
        while (sc.hasNextLine()) {
            lines.add(sc.nextLine());
        }

        char[] chars = new char[12];
        for (int i = 0; i < chars.length; i++) {
            int z = 0;
            int o = 0;
            for (String input : lines) {
                if (input.charAt(i) == '1') {
                    o += 1;
                    continue;
                }

                z += 1;
            }

            if (z > o) {
                chars[i] = '0';
                continue;
            }

            chars[i] = '1';
        }

        String binary = new String(chars);
        String inverted = String.join("", Arrays.stream(binary.split("")).map(m -> {
            if (m.equals("0")) return "1";
            if (m.equals("1")) return "0";
            return "";
        }).collect(Collectors.toList()));
        int gamma = Integer.parseInt(binary, 2);
        int epsilon = Integer.parseInt(inverted, 2);
        
        System.out.println("Binary   : " + binary);
        System.out.println("Gamme    : " + gamma);
        System.out.println("Epsilon  : " + epsilon);
        System.out.println("PowerCons: " + epsilon * gamma);
        
    }
}