PORTNAME=	skopeo
PORTVERSION=	1.22.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for various operations on container images and repositories
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	containers

PLIST_FILES=	bin/${PORTNAME}

#CGO_ENABLED=0 -> imports github.com/proglottis/gpgme: build constraints exclude all Go files
#-extldflags -static -> ld: error: unable to find library -lgpgme
GO_BUILDFLAGS+=	-ldflags="-s -w"

# vendor directory appears to be pre-populated
post-patch: .PHONY
	@${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@${TEST} -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@${TEST} -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/go.podman.io
	@${MKDIR} ${WRKSRC}/vendor/go.podman.io/image
	@${TEST} -z ${WRKSRC_containers_container_libs_common} || ${RLN} ${WRKSRC_containers_container_libs_common}/common ${WRKSRC}/vendor/go.podman.io
	@${TEST} -z ${WRKSRC_containers_container_libs_storage} || ${RLN} ${WRKSRC_containers_container_libs_storage}/storage ${WRKSRC}/vendor/go.podman.io
	@${TEST} -z ${WRKSRC_containers_container_libs_image} || ${RLN} ${WRKSRC_containers_container_libs_image}/image ${WRKSRC}/vendor/go.podman.io/image/v5

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-moby:moby:v0.8.2:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.8.2:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/)(docker-credential-helpers)$|docker:\2\1\2|1}

#-open-telemetry:opentelemetry-go-contrib:v0.45.0:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
#-open-telemetry:opentelemetry-go-instrumentation:v1.1.0:open_telemetry_opentelemetry_go_instrumentation/vendor/go.opentelemetry.io/auto/sdk
#-open-telemetry:opentelemetry-proto-go:v1.0.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto/otlp
GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:opentelemetry-go*}

#-google:go-genproto:207652e42e2e:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:207652e42e2e:google_go_genproto_api
#-google:go-genproto:207652e42e2e:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:207652e42e2e:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C,_genproto.*/vendor/google\.golang\.org/genproto/googleapis/(api|rpc),_genproto_\1,}

#-containers:container-libs:v0.66.0:containers_container_libs/vendor/go.podman.io/common
#+containers:container-libs:common/v0.66.0:containers_container_libs_common
#-containers:container-libs:v1.61.0:containers_container_libs_2/vendor/go.podman.io/storage
#+containers:container-libs:storage/v1.61.0:containers_container_libs_storage
#-containers:container-libs:v5.38.0:containers_container_libs_1/vendor/go.podman.io/image/v5
#+containers:container-libs:image/v5.38.0:containers_container_libs_image
GH_TUPLE:=	${GH_TUPLE:C,^(containers):(container)-(libs):([^:]+).*/go\.podman\.io/(common|storage|image).*,\1:\2-\3:\5/\4:\1_\2_\3_\5,}

.include <bsd.port.mk>
