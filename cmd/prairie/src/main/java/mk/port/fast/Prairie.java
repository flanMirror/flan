package mk.port.fast;

import de.neuland.pug4j.Pug4J;
import de.neuland.pug4j.template.PugTemplate;
import org.apache.commons.io.FileUtils;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Objects;

public class Prairie {
    public static String in;
    public static String out;

    public static void main(String[] args) {
        if (args.length < 2) {
            System.out.println("need input, output and optional wipe");
            System.exit(1);
        }

        in = args[0] + "/packages/backend/src/server/web/views/";
        out = args[1];
        de.neuland.pug4j.parser.PathHelper.in = in;

        File o = new File(out);
        if (o.exists()) {
            if (args.length == 3 && Objects.equals(args[2], "wipe")) {
                try {
                    FileUtils.deleteDirectory(o);
                } catch (IOException e) {
                    System.out.printf("error wiping directory: %s", e.getMessage());
                    System.exit(1);
                }
            } else {
                System.out.println("output directory not empty, refusing to continue");
                System.exit(1);
            }
        }
        if (!o.mkdir()) {
            System.out.println("wasn't able to create output directory");
            System.exit(1);
        }

        File[] targets = FileUtils.listFiles(new File(in), new String[]{"pug"}, false).toArray(new File[]{});
        try {
            for (File target : targets) {
                String name = target.getName();
                PugTemplate template = Pug4J.getTemplate(in + name);
                name = name.substring(0, name.length() - 3) + ".tmpl";
                System.out.printf("generating %s", name);
                long t = System.nanoTime();
                Pug4J.render(template, null, new FileWriter(o.getAbsolutePath() + "/" + name));
                System.out.printf(" (%.2fms)\n", (System.nanoTime() - t) / 1000000.0);
            }
        } catch (IOException e) {
            System.out.printf("error reading input file: %s\n", e.getMessage());
            System.exit(1);
        }
    }
}
