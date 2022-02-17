package de.neuland.pug4j.parser;

import de.neuland.pug4j.exceptions.PugTemplateLoaderException;
import org.apache.commons.io.FilenameUtils;

import java.util.Objects;

public class PathHelper {
    public static String in;

    public String resolvePath(String parentTemplateName, String templateName, String basePath) {
        parentTemplateName = FilenameUtils.separatorsToUnix(parentTemplateName);
        templateName = FilenameUtils.separatorsToUnix(templateName);

        if (FilenameUtils.getPrefixLength(basePath) != 0)
            throw new PugTemplateLoaderException("basePath " + basePath + " must be relative");

        String inputParamsLog = "ParentFilename: " + parentTemplateName + ", TemplateName: " + templateName + ", BasePath:" + basePath;

        if (parentTemplateName.startsWith("/") && basePath.length() > 0) {
            parentTemplateName = FilenameUtils.normalize(basePath + parentTemplateName, true);
        }

        if (templateName.startsWith("/") && basePath.length() > 0) {
            String path = FilenameUtils.normalize(basePath + templateName, true);
            return path;
        }

        String parent = FilenameUtils.getPath(parentTemplateName);

        // begin ugliness
        if (Objects.equals(parent, "./")) {
            return templateName;
        }

        if (Objects.equals(parent, "")) {
            if (templateName.endsWith(".pug")) {
                return templateName;
            }
            return FilenameUtils.normalize(FilenameUtils.concat(in, templateName), true);
        }
        // end ugliness

        if (parent == null) {
            return templateName;
        }

        String resolve = FilenameUtils.normalize(FilenameUtils.concat(parent, templateName), true);
        System.out.println(resolve);
        return resolve;
    }
}
