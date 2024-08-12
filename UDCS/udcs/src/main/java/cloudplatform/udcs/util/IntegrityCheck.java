package cloudplatform.udcs.util;

import jakarta.annotation.PostConstruct;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.util.Arrays;

@Component
public class IntegrityCheck {

    @Value("${log.path}")
    private String logPath;

    static final String startPattern = "Type = ENROLL";
    static final String endPattern = "Type = END";

    @PostConstruct
    public void integrityCheck(){

        File logDirectory = new File(logPath);
        File[] logFiles = logDirectory.listFiles((dir, name) -> name.matches("\\d{4}-\\d{2}-\\d{2}\\.log"));
        System.out.println("현재 경로");
        System.out.println(logDirectory.getAbsolutePath());
        if (logFiles != null && logFiles.length > 0) {
            Arrays.sort(logFiles, (f1, f2) -> {
                LocalDate date1 = LocalDate.parse(f1.getName().substring(0, 10), DateTimeFormatter.ofPattern("yyyy-MM-dd"));
                LocalDate date2 = LocalDate.parse(f2.getName().substring(0, 10), DateTimeFormatter.ofPattern("yyyy-MM-dd"));
                return date2.compareTo(date1);
            });

            // 최신 로그 파일 출력
            File latestLogFile = logFiles[0];
            System.out.println("최신 로그 파일: " + latestLogFile.getName());
        } else {
            System.out.println("로그 파일이 존재하지 않습니다.");
        }

        boolean inProgress = false;

        try (BufferedReader br = new BufferedReader(new FileReader(logPath))) {
            String line;
            while ((line = br.readLine()) != null) {
                if (!inProgress && line.contains(startPattern)) {
                    System.out.println("이벤트 시작:");
                    System.out.println(line);
                } else if (inProgress && line.contains(endPattern)) {
                    System.out.println(line);
                    inProgress = false;
                    System.out.println("이벤트 종료");
                } else if (inProgress) {
                    System.out.println(line);
                }
            }

            if (inProgress) {
                System.out.println("이벤트가 비정상적으로 종료되었습니다.");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

}
