package demo.bank.customer;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@Service
public class ExternalCustomerRepository {
    private static final Logger logger = LoggerFactory.getLogger(ExternalCustomerRepository.class);

    private RestTemplateBuilder restTemplateBuilder;
    static private ThreadLocal<RestTemplate> holder = new ThreadLocal<>();

    @Value("${external.customer.url}")
    private String extCustomerSvcUrl;

    @Autowired
    public ExternalCustomerRepository(RestTemplateBuilder restTemplateBuilder) {
        this.restTemplateBuilder = restTemplateBuilder;
    }

    public Map<String, String> getExternalCustomerById(String customerId) {
        logger.info("external.customer.url={}", extCustomerSvcUrl);
        Map<String, String> response = getLocalRestTemplate().getForObject(extCustomerSvcUrl + "/" + customerId, Map.class);
        return response;
    }

    private RestTemplate getLocalRestTemplate() {
        RestTemplate template = holder.get();
        if (template == null) {
            template = restTemplateBuilder.build();
            holder.set(template);
        }
        return template;
    }
}
