PORTNAME=	docker
PORTVERSION=	29.9.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Open-source application container engine
WWW=		https://www.docker.com/

LICENSE=	APACHE20

RUN_DEPENDS=	git:devel/git

USES=		cpe go

USE_GITHUB=	yes
GH_ACCOUNT=	docker
GH_PROJECT=	cli

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X docker/vendor/github.com/${GH_ACCOUNT}/${GH_PROJECT}/cli/version.BuildTime=${DATE_EPOCH_CMD:sh}
GO_LDFLAGS+=	-X docker/vendor/github.com/${GH_ACCOUNT}/${GH_PROJECT}/cli/version.GitCommit=${COMMIT_CMD:sh}
GO_LDFLAGS+=	-X docker/vendor/github.com/${GH_ACCOUNT}/${GH_PROJECT}/cli/version.PlatformName=${OPSYS}/${ARCH}
GO_LDFLAGS+=	-X docker/vendor/github.com/${GH_ACCOUNT}/${GH_PROJECT}/cli/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/docker

post-extract: .PHONY
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${MKDIR} ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/cli ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/cli-plugins ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/cmd ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/internal ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/opts ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/pkg ${WRKSRC}/vendor/github.com/docker/cli
	${RLN} ${WRKSRC}/templates ${WRKSRC}/vendor/github.com/docker/cli

.include <bsd.port.pre.mk>

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL}^{} | cut -c1-7

DATE_EPOCH_CMD=	date -ur \
		$$(${GREP} --max-count=1 TIMESTAMP ${DISTINFO_FILE} | ${CUT} -d' ' -f3) \
		'+%FT%TZ'

.include <bsd.port.post.mk>
