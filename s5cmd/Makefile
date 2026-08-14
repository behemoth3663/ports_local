PORTNAME=	s5cmd
PORTVERSION=	2.3.0
DISTVERSIONPREFIX=	v
PORTREVISION=	20
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Parallel S3 and local filesystem execution tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	peak

PLIST_FILES=	bin/${PORTNAME}

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL} | cut -c1-7

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/version.Version=${DISTVERSIONFULL} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/version.GitCommit=$$(${COMMIT_CMD})"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-cheggaaa:pb:v3.1.4:cheggaaa_pb_v3/vendor/github.com/cheggaaa/pb
#+cheggaaa:pb:v3.1.4:cheggaaa_pb_v3/vendor/github.com/cheggaaa/pb/v3
GH_TUPLE:=	${GH_TUPLE:S|/cheggaaa/pb/v3|/cheggaaa/pb|1}

.include <bsd.port.mk>
