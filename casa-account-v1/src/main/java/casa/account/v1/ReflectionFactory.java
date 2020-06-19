package casa.account.v1;

import io.grpc.BindableService;
import io.grpc.protobuf.services.ProtoReflectionService;
import io.micronaut.context.annotation.Factory;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Singleton;

@Factory
class ReflectionFactory {
  private static final Logger logger = LoggerFactory.getLogger(ReflectionFactory.class);

  @Singleton
  BindableService reflectionService() {
    logger.info("initializing gRPC reflection service");
    return ProtoReflectionService.newInstance();
  }
}
