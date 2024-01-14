package com.example.config;

import com.google.gson.Gson;
import lombok.Getter;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

//import javax.annotation.PostConstruct;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Locale;
import java.util.concurrent.TimeUnit;

@SuppressWarnings("SpellCheckingInspection")
@Component
@Profile("default")
@Getter
public class BaseConfig {
    protected String PROFILE_NAME = "default";

    // App Setting
    protected String VERSION_CODE = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyyMMddHHmmss"));
    protected String APP_NAME = "example";

    // Locale
    protected Locale[] SUPPORTED_LOCALE = {Locale.US, Locale.KOREA};

    // CACHE
    protected long CACHE_DURATION_TIME = 60;
    protected TimeUnit CACHE_DURATION_TIME_UNIT = TimeUnit.MINUTES;

    public void initialize() {
        System.out.print(this.toString());
    }

    @Override
    public String toString() {
        return new Gson().toJson(this);
    }
}
