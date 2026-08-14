PORTNAME=	s5cmd
PORTVERSION=	2.3.0
DISTVERSIONPREFIX=	v
PORTREVISION=	22
CATEGORIES=	local net
PKGNAMESUFFIX=	-local

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

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_
	cd ${WRKSRC} && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${PRINTF} '//go:build tools\npackage tools\nimport _ "github.com/rivo/uniseg"\n' >require.go && \
		${GO_CMD} mod edit -replace=github.com/rivo/uniseg=github.com/rivo/uniseg@v0.2.0 && \
		${GO_CMD} mod edit -replace=github.com/lanrat/extsort=github.com/lanrat/extsort@v1.0.3 && \
		${GO_CMD} get -u && \
		${GO_CMD} mod tidy && \
		${SETENV} TZ=UTC ${DIFF} -u --minimal -- go.mod.orig go.mod | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}_/patch-go.mod && \
		${SETENV} TZ=UTC ${DIFF} -u --minimal -- go.sum.orig go.sum | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}_/patch-go.sum

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
