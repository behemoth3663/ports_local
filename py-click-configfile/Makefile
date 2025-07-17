PORTNAME=	click-configfile
PORTVERSION=    0.2.3
CATEGORIES=	local devel python
MASTER_SITES=	PYPI
PKGNAMEPREFIX=	${PYTHON_PKGNAMEPREFIX}

MAINTAINER=	none@freebsd.org
COMMENT=	Click commands that use configuration files
WWW=		https://github.com/click-contrib/click-configfile

LICENSE=	BSD3CLAUSE

RUN_DEPENDS=	${PYTHON_PKGNAMEPREFIX}click>=6.6:devel/py-click@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}six>=1.15:devel/py-six@${PY_FLAVOR} \
		${PYTHON_PKGNAMEPREFIX}configparser>=3.5.0:devel/py-configparser@${PY_FLAVOR}

USES=		python
USE_PYTHON=	autoplist concurrent pep517

.include <bsd.port.mk>
