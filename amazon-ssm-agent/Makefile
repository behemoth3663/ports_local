#

PORTNAME=	amazon-ssm-agent
PORTVERSION=	3.1.1856.0
PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	cperciva@FreeBSD.org
COMMENT=	Amazon Simple Systems Manager Agent
WWW=		https://github.com/aws/amazon-ssm-agent

LICENSE=	APACHE20

RUN_DEPENDS=	${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_SUBDIR=	src/github.com/aws/amazon-ssm-agent

JQ_GETVER_EXPR=	[.[].name | select(test("^[-0-9.]+$$"; "n"))][0]

GO_TARGET=	./agent/framework/processor/executer/outofproc/worker:${PREFIX}/sbin/ssm-document-worker \
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

post-patch: .PHONY
	${REINPLACE_CMD} -r -e \
		's|(const[[:space:]]+Version[=[:space:]]+)"[.0-9]+"|\1"${PORTVERSION}"|' \
	${WRKSRC}/agent/version/version.go

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${PREFIX}/etc/amazon/ssm
	${INSTALL_DATA} ${WRKSRC}/amazon-ssm-agent.json.template ${STAGEDIR}${PREFIX}/etc/amazon/ssm
	${INSTALL_DATA} ${WRKSRC}/seelog_unix.xml ${STAGEDIR}${PREFIX}/etc/amazon/ssm/seelog_unix.xml.template

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
