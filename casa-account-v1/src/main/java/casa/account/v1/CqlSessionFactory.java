package casa.account.v1;

import com.datastax.oss.driver.api.core.CqlSession;
import com.datastax.oss.driver.api.core.cql.ResultSet;
import io.micronaut.context.annotation.Bean;
import io.micronaut.context.annotation.Factory;
import io.micronaut.context.annotation.Value;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Singleton;
import java.net.InetSocketAddress;
import java.nio.file.Files;
import java.nio.file.Paths;

@Factory
public class CqlSessionFactory {
  private static final Logger logger = LoggerFactory.getLogger(CqlSessionFactory.class);

  @Value("${cassandra.instance:astra}")
  private String cassandraInstance;

  @Value("${cassandra.username:vino9}")
  private String dbUsername;

  @Value("${cassandra.password:secret}")
  private String dbPassword;

  @Value("${cassandra.host:localhost}")
  private String localCassandraHost;

  @Value("${cassandra.port:9042}")
  private int localCassandraPort;

  private CqlSession session;

  @Bean(preDestroy = "close")
  @Singleton
  public CqlSession initializeCassandraSession() {
    logger.info(
        "cassandra config instance={}, host={}, port={}, username={}, password={}****",
        cassandraInstance,
        localCassandraHost,
        localCassandraPort,
        dbUsername,
        dbPassword.substring(0, 1));

    String cwd = Paths.get(".").toAbsolutePath().normalize().toString();
    logger.info("CWD={}", cwd);

    // the secure bundle is in the parent directory when running locally outside
    // K8S environment.
    // when running inside K8s it is at the /astra directory mounted from a secret
    String secureBundle = "../secure-connect-vino9.zip";
    if (!Files.exists(Paths.get(secureBundle).toAbsolutePath())) {
      secureBundle = "/astra/secure-connect-vino9.zip";
    }

    if ("astra".equalsIgnoreCase(cassandraInstance)) {
      session =
          CqlSession.builder()
              .withCloudSecureConnectBundle(Paths.get(secureBundle))
              .withAuthCredentials(dbUsername, dbPassword)
              .withKeyspace("vino9")
              .build();
    } else if ("local".equalsIgnoreCase(cassandraInstance)) {
      session =
          CqlSession.builder()
              .addContactPoint(new InetSocketAddress(localCassandraHost, localCassandraPort))
              .withAuthCredentials(dbUsername, dbPassword)
              .withKeyspace("vino9")
              .withLocalDatacenter("Cassandra")
              .build();
    } else {
      logger.info("cassandra not available");
      return null;
    }

    ResultSet rs = session.execute("select release_version from system.local");
    logger.info("connected to cassandra version {}", rs.one().getString("release_version"));

    return session;
  }
}
