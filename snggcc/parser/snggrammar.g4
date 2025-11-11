grammar snggrammar;

NEWLINE : [\r\n] ;
NUMBER  : [0-9] ;
WS: [ \t];
CHAR: [a-zA-Z];
UNDERSCORE: '_';
MINUS: '-';
ESCAPE_PREFIX: '\\';
TEXTLINE: [a-zA-Z];
DQUOTES: '"';
PUNCTUATION: [,.:;!?'`];

script:   (scriptpart NEWLINE NEWLINE)* ;
id: CHAR (CHAR | NUMBER)+;
escapesequence: ESCAPE_PREFIX (ESCAPE_PREFIX | DQUOTES);
string: DQUOTES (CHAR | escapesequence | WS | PUNCTUATION)+ DQUOTES;
scriptpart: partheader (headeroption)+ NEWLINE partbody;
partheader: parttype WS+ partid WS+ partname WS* NEWLINE;
headeroption: optionid ':' string NEWLINE;
parttype: id;
partid: id;
optionid: id;
partname: string;
partbody: multiline;
text: (word worddelim)+;
multiline: (text NEWLINE{0,1})+;
word: (CHAR | MINUS | UNDERSCORE | NUMBER)+;
worddelim: (WS|PUNCTUATION|DQUOTES)+;
