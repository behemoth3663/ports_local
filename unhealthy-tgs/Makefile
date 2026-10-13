PORTNAME=	unhealthy-tgs
PORTVERSION=	0.1.37
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Get unhealthy target groups of Load Balancers
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}/tree/main/cmd/${PORTNAME}

LICENSE=	BSD3CLAUSE

RUN_DEPENDS=	garble:local/garble

USES=		go:1.27,modules

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi
GH_PROJECT=	go

WRKSRC_SUBDIR=	cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X main.version=${PORTVERSION}
GO_LDFLAGS+=	-X main.commit=${COMMIT_CMD:sh}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	${RM} ${WRKSRC}/../../go.work
	${MKDIR} ${WRKSRC}/vendor
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

do-build: .PHONY
	( cd ${GO_WRKSRC} && ${ECHO_MSG} "===>  Building ${PORTNAME} from ." && \
		${SETENVI} ${WRK_ENV} ${MAKE_ENV} ${GO_ENV} GOMAXPROCS=${MAKE_JOBS_NUMBER} GOPROXY=off \
			${LOCALBASE}/bin/garble -tiny build ${GO_BUILDFLAGS} -o ${GO_WRKDIR_BIN}/${PORTNAME} . )

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

.include <bsd.port.pre.mk>

COMMIT_CMD=	git ls-remote \
	https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
	refs/tags/${DISTVERSIONFULL}^{} | cut -c1-7

.include <bsd.port.post.mk>
