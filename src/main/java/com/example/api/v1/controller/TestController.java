package com.example.api.v1.controller;

import com.example.config.BaseConfig;
import com.example.controller.CoreController;
import com.example.model.msg.Message;
import com.example.model.msg.MessageStatus;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.ExampleObject;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.i18n.LocaleContextHolder;
import org.springframework.web.bind.annotation.*;

import java.io.UnsupportedEncodingException;
import java.util.HashMap;
import java.util.Locale;
import java.util.Map;

@RestController
@RestControllerAdvice
@RequestMapping("/api/v1/test")
@Tag(name = "Test controller")
public class TestController extends CoreController {

    @Autowired
    public TestController(BaseConfig configData) {
        super(configData);
    }

    @GetMapping("")
    @ResponseBody
    @Operation(
            summary = "/",
            description = "some description"
    )
    @ApiResponse(
            responseCode = "200",
            description = "some description",
            content = {
                    @Content(
                            examples = {
                                    @ExampleObject(name = "/",
                                            summary = "some summary",
                                            value = "{\"status\":200,\"message\":\"처리가 완료되었습니다.\",\"ts\":\"20240114113903\",\"response\":{\"Hello\":\"world\"}}"
                                    )
                            },
                            mediaType = "application/json", schema = @Schema(implementation = Message.class))
            }
    )

    public Message greeter(HttpServletRequest request) throws UnsupportedEncodingException {
        Map<String, Object> result = new HashMap<>();
        result.put("Hello", "world");

        // in header
        // Accept-Language: en-US
        // Accept-Language: ko
        Locale currentLocale = LocaleContextHolder.getLocale();
        return new Message(MessageStatus.RES_HTTP_OK, result, currentLocale);
    }

}
