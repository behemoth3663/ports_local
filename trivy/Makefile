PORTNAME=	trivy
PORTVERSION=	0.75.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local security
PKGNAMESUFFIX=	-local

MAINTAINER=	mfechner@FreeBSD.org
COMMENT=	Security scanner written in go
WWW=		https://trivy.dev https://${GO_MODULE}

LICENSE=	APACHE20

ONLY_FOR_ARCHS=		aarch64 amd64
ONLY_FOR_ARCHS_REASON=	module modernc.org/sqlite only supported on aarch64 and amd64

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aquasecurity
GO_MOD_DIST=	github

GO_MODULE=	github.com/${GH_ACCOUNT}/${GH_PROJECT}
GO_TARGET=	./cmd/${PORTNAME}

#CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-extldflags -static
GO_LDFLAGS+=	-X ${GO_MODULE}/pkg/version/app.ver=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=jsonv2,norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/github.com/GoogleCloudPlatform/opentelemetry-operations-go
	${RM} -r ${WRKSRC}/vendor/github.com/Microsoft/hcsshim
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/experimental/opentelemetry
	${RM} -r ${WRKSRC}/vendor/google.golang.org/grpc/stats/opentelemetry
	${RM} -r ${WRKSRC}/vendor/otelconnect

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${DATADIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/*.tpl ${STAGEDIR}${DATADIR}/
	${FIND} -s ${STAGEDIR}${DATADIR} -not -type d | ${SORT} | \
		${SED} -e 's#^${STAGEDIR}${PREFIX}/##' | \
		${SED} -E -e '/sample$$/ s#^#@sample #; \
		s#${DATADIR_REL}/bin#@(,,555) ${DATADIR_REL}/bin#; ' >> ${TMPPLIST}
	${CAT} ${TMPPLIST}

.include <bsd.port.mk>
