package com.example.controller;

import com.example.config.BaseConfig;
import com.example.exception.Forbidden;
import com.example.exception.UnAuthorized;
import com.example.model.msg.Message;
import com.example.model.msg.MessageStatus;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.NoHandlerFoundException;

@SuppressWarnings("SpellCheckingInspection")
@RestController
@RestControllerAdvice
public abstract class CoreController {
    protected final BaseConfig configData;

    @Controller
    static class FaviconController {
        @GetMapping("favicon.ico")
        @ResponseBody
        void returnNoFavicon() {
        }
    }
    @Autowired
    protected CoreController(BaseConfig configData) {
        this.configData = configData;
    }

    /* 400 */
    @ExceptionHandler(MissingServletRequestParameterException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ResponseBody
    public Message handleParamMissError(MissingServletRequestParameterException e) {
        return new Message(MessageStatus.RES_HTTP_BAD_REQUEST);
    }

    /* 401 */
    @ExceptionHandler(UnAuthorized.class)
    @ResponseStatus(HttpStatus.UNAUTHORIZED)
    @ResponseBody
    public Message handleUnAuthorizedError(UnAuthorized e) {
        return new Message(MessageStatus.RES_HTTP_UNAUTHORIZED);
    }

    /* 403 */
    @ExceptionHandler(Forbidden.class)
    @ResponseStatus(HttpStatus.FORBIDDEN)
    @ResponseBody
    public Message handleForbiddenError(Exception e) {
        return new Message(MessageStatus.RES_HTTP_FORBIDDEN);
    }


    /* 404 */
    @ExceptionHandler(NoHandlerFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    @ResponseBody
    public Message handleNotFoundPageError(Exception e) {
        return new Message(MessageStatus.RES_HTTP_NOT_FOUND);
    }

    /* 500 */
    @ExceptionHandler(Exception.class)
    @ResponseBody
    public Message handleServerError(Exception e) {
        e.printStackTrace();
        return new Message(MessageStatus.RES_HTTP_SERVER_ERROR);
    }
}