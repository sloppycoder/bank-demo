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

    @Value("${external.customer.url}")
    private String extCustomerSvcUrl;

    @Autowired
    public ExternalCustomerRepository(RestTemplateBuilder restTemplateBuilder) {
        this.restTemplateBuilder = restTemplateBuilder;
    }

    public Map<String, String> getExternalCustomerById(String customerId) {
        logger.info("external.customer.url={}", extCustomerSvcUrl);
        RestTemplate template = restTemplateBuilder.build();
        Map<String, String> response = template.getForObject(extCustomerSvcUrl + "/" + customerId, Map.class);
        return response;
    }
}
