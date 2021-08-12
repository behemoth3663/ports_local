#

PORTNAME=	amazon-ssm-agent
PORTVERSION=	3.1.192.0
CATEGORIES=	sysutils

MAINTAINER=	cperciva@FreeBSD.org
COMMENT=	Amazon Simple Systems Manager Agent

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_SUBDIR=	src/github.com/aws/amazon-ssm-agent

JQ_GETVER_EXPR=	[.[].name | select(test("^[-0-9.]+$$"; "n"))][0]

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
#.-include "${.CURDIR}/Makefile.deps"
#post-patch-A1VM-on: .PHONY
#	${REINPLACE_CMD} -r -e \
#		's|(const[[:space:]]+Version[=[:space:]]+)"[.0-9]+"|\1"${PORTVERSION}"|' \
#	${WRKSRC}/agent/version/version.go
#	${FIND} ${WRKSRC}/src ${WRKSRC}/vendor -type f -name '*.go' \
#		-exec ${REINPLACE_CMD} -r 's,"github\.com/go-git/(go-git|go-billy)/v5,"github.com/go-git/\1,' {} \+

GO_BUILDFLAGS=	-ldflags "-s -w"

GO_TARGET=	\
		./agent/framework/processor/executer/outofproc/worker:${PREFIX}/sbin/ssm-document-worker \
		./agent/framework/processor/executer/outofproc/sessionworker:${PREFIX}/sbin/ssm-session-worker \
		./agent/session/logging:${PREFIX}/sbin/ssm-session-logger \
		./agent:${PREFIX}/sbin/ssm-agent-worker \
		./agent/cli-main:${PREFIX}/sbin/ssm-cli \
		./core:${PREFIX}/sbin/amazon-ssm-agent

PLIST_FILES=	etc/amazon/ssm/amazon-ssm-agent.json.template	\
		etc/amazon/ssm/seelog_unix.xml.template		\
		sbin/ssm-document-worker			\
		sbin/ssm-session-worker				\
		sbin/ssm-session-logger				\
		sbin/ssm-agent-worker				\
		sbin/ssm-cli					\
		sbin/amazon-ssm-agent

USE_RC_SUBR=	amazon-ssm-agent

EXTRA_PATCHES+=	${FILESDIR}/2fce6aca986ea92277f4726c7acae915d15054ed.patch:-p1
EXTRA_PATCHES+=	${FILESDIR}/b084979a353b90bbec8d60d27858bd9c3067ee1c.patch:-p1

#post-patch:
#	${LN} -s ${WRKSRC}/vendor/src/* ${WRKSRC}/vendor/

#pre-build:
# set package version
#	cd ${GO_WRKSRC} && BRAZIL_PACKAGE_VERSION=${PORTVERSION} \
#	${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} run ./agent/version/versiongenerator/version-gen.go

post-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/etc/amazon/ssm
	${INSTALL_DATA} ${WRKSRC}/amazon-ssm-agent.json.template ${STAGEDIR}${PREFIX}/etc/amazon/ssm
	${INSTALL_DATA} ${WRKSRC}/seelog_unix.xml ${STAGEDIR}${PREFIX}/etc/amazon/ssm/seelog_unix.xml.template

.include <bsd.port.mk>
