// Generated from /Users/kristofer/LocalProjects/Unified/src/unified-compiler/grammar/UnifiedParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class UnifiedParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LBRACE=1, RBRACE=2, LPAREN=3, RPAREN=4, LBRACK=5, RBRACK=6, SEMI=7, COLON=8, 
		COMMA=9, DOT=10, ARROW=11, DOUBLE_COLON=12, PLUS=13, MINUS=14, STAR=15, 
		DIV=16, MOD=17, BIT_AND=18, BIT_OR=19, BIT_XOR=20, BIT_NOT=21, LSHIFT=22, 
		RSHIFT=23, AND=24, OR=25, NOT=26, EQ=27, NE=28, LT=29, GT=30, LE=31, GE=32, 
		ASSIGN=33, PLUS_ASSIGN=34, MINUS_ASSIGN=35, STAR_ASSIGN=36, DIV_ASSIGN=37, 
		MOD_ASSIGN=38, LSHIFT_ASSIGN=39, RSHIFT_ASSIGN=40, AND_ASSIGN=41, XOR_ASSIGN=42, 
		OR_ASSIGN=43, QUESTION=44, QUESTION_QUESTION=45, NULL_COND=46, RANGE=47, 
		RANGE_INCL=48, INC=49, DEC=50, UNDERSCORE=51, MODULE=52, IMPORT=53, AS=54, 
		FN=55, STRUCT=56, ENUM=57, INTERFACE=58, IMPL=59, ACTOR=60, TYPE=61, REFINE=62, 
		CONST=63, LET=64, VAR=65, MUT=66, PUB=67, SELF=68, IF=69, ELSE=70, LOOP=71, 
		WHILE=72, FOR=73, IN=74, BREAK=75, CONTINUE=76, RETURN=77, SWITCH=78, 
		CASE=79, WHERE=80, MOVE=81, AWAIT=82, ASYNC=83, TRY=84, REGION=85, BoolLiteral=86, 
		NullLiteral=87, Identifier=88, IntLiteral=89, FloatLiteral=90, StringLiteral=91, 
		CharLiteral=92, EscapeSequence=93, WS=94, COMMENT=95, MULTILINE_COMMENT=96;
	public static final int
		RULE_program = 0, RULE_item = 1, RULE_moduleDecl = 2, RULE_importDecl = 3, 
		RULE_importPath = 4, RULE_importList = 5, RULE_functionDecl = 6, RULE_paramList = 7, 
		RULE_parameter = 8, RULE_genericParams = 9, RULE_genericParam = 10, RULE_typeConstraint = 11, 
		RULE_whereClause = 12, RULE_whereConstraint = 13, RULE_structDecl = 14, 
		RULE_structMember = 15, RULE_enumDecl = 16, RULE_enumVariant = 17, RULE_enumVariantParams = 18, 
		RULE_enumVariantParam = 19, RULE_interfaceDecl = 20, RULE_interfaceMember = 21, 
		RULE_functionSig = 22, RULE_implDecl = 23, RULE_implMember = 24, RULE_actorDecl = 25, 
		RULE_actorMember = 26, RULE_typeAlias = 27, RULE_constantDecl = 28, RULE_type = 29, 
		RULE_typeList = 30, RULE_statement = 31, RULE_letStatement = 32, RULE_varStatement = 33, 
		RULE_assignmentStatement = 34, RULE_assignmentOp = 35, RULE_regionStatement = 36, 
		RULE_exprStatement = 37, RULE_returnStatement = 38, RULE_ifStatement = 39, 
		RULE_loopStatement = 40, RULE_whileStatement = 41, RULE_forStatement = 42, 
		RULE_switchStatement = 43, RULE_caseClause = 44, RULE_breakStatement = 45, 
		RULE_continueStatement = 46, RULE_blockStatement = 47, RULE_pattern = 48, 
		RULE_patternList = 49, RULE_fieldPattern = 50, RULE_fieldPatternList = 51, 
		RULE_expr = 52, RULE_caseExpr = 53, RULE_primary = 54, RULE_lambdaExpr = 55, 
		RULE_asyncExpr = 56, RULE_structExpr = 57, RULE_fieldInitList = 58, RULE_fieldInit = 59, 
		RULE_listExpr = 60, RULE_mapExpr = 61, RULE_keyValue = 62, RULE_setExpr = 63, 
		RULE_tupleExpr = 64, RULE_namedTupleField = 65, RULE_block = 66, RULE_argList = 67, 
		RULE_literal = 68, RULE_identifier = 69;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "item", "moduleDecl", "importDecl", "importPath", "importList", 
			"functionDecl", "paramList", "parameter", "genericParams", "genericParam", 
			"typeConstraint", "whereClause", "whereConstraint", "structDecl", "structMember", 
			"enumDecl", "enumVariant", "enumVariantParams", "enumVariantParam", "interfaceDecl", 
			"interfaceMember", "functionSig", "implDecl", "implMember", "actorDecl", 
			"actorMember", "typeAlias", "constantDecl", "type", "typeList", "statement", 
			"letStatement", "varStatement", "assignmentStatement", "assignmentOp", 
			"regionStatement", "exprStatement", "returnStatement", "ifStatement", 
			"loopStatement", "whileStatement", "forStatement", "switchStatement", 
			"caseClause", "breakStatement", "continueStatement", "blockStatement", 
			"pattern", "patternList", "fieldPattern", "fieldPatternList", "expr", 
			"caseExpr", "primary", "lambdaExpr", "asyncExpr", "structExpr", "fieldInitList", 
			"fieldInit", "listExpr", "mapExpr", "keyValue", "setExpr", "tupleExpr", 
			"namedTupleField", "block", "argList", "literal", "identifier"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "'('", "')'", "'['", "']'", "';'", "':'", "','", 
			"'.'", "'->'", "'::'", "'+'", "'-'", "'*'", "'/'", "'%'", "'&'", "'|'", 
			"'^'", "'~'", "'<<'", "'>>'", "'&&'", "'||'", "'!'", "'=='", "'!='", 
			"'<'", "'>'", "'<='", "'>='", "'='", "'+='", "'-='", "'*='", "'/='", 
			"'%='", "'<<='", "'>>='", "'&='", "'^='", "'|='", "'?'", "'??'", "'?.'", 
			"'..'", "'..='", "'++'", "'--'", "'_'", null, "'import'", "'as'", "'fn'", 
			"'struct'", "'enum'", "'interface'", "'impl'", "'actor'", "'type'", "'refine'", 
			"'const'", "'let'", "'var'", "'mut'", "'pub'", null, "'if'", "'else'", 
			"'loop'", "'while'", "'for'", "'in'", "'break'", "'continue'", "'return'", 
			"'switch'", "'case'", "'where'", "'move'", "'await'", "'async'", "'try'", 
			"'region'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LBRACE", "RBRACE", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "SEMI", 
			"COLON", "COMMA", "DOT", "ARROW", "DOUBLE_COLON", "PLUS", "MINUS", "STAR", 
			"DIV", "MOD", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_NOT", "LSHIFT", "RSHIFT", 
			"AND", "OR", "NOT", "EQ", "NE", "LT", "GT", "LE", "GE", "ASSIGN", "PLUS_ASSIGN", 
			"MINUS_ASSIGN", "STAR_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN", 
			"RSHIFT_ASSIGN", "AND_ASSIGN", "XOR_ASSIGN", "OR_ASSIGN", "QUESTION", 
			"QUESTION_QUESTION", "NULL_COND", "RANGE", "RANGE_INCL", "INC", "DEC", 
			"UNDERSCORE", "MODULE", "IMPORT", "AS", "FN", "STRUCT", "ENUM", "INTERFACE", 
			"IMPL", "ACTOR", "TYPE", "REFINE", "CONST", "LET", "VAR", "MUT", "PUB", 
			"SELF", "IF", "ELSE", "LOOP", "WHILE", "FOR", "IN", "BREAK", "CONTINUE", 
			"RETURN", "SWITCH", "CASE", "WHERE", "MOVE", "AWAIT", "ASYNC", "TRY", 
			"REGION", "BoolLiteral", "NullLiteral", "Identifier", "IntLiteral", "FloatLiteral", 
			"StringLiteral", "CharLiteral", "EscapeSequence", "WS", "COMMENT", "MULTILINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "UnifiedParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public UnifiedParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<ItemContext> item() {
			return getRuleContexts(ItemContext.class);
		}
		public ItemContext item(int i) {
			return getRuleContext(ItemContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 52)) & ~0x3f) == 0 && ((1L << (_la - 52)) & 35835L) != 0)) {
				{
				{
				setState(140);
				item();
				}
				}
				setState(145);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ItemContext extends ParserRuleContext {
		public ModuleDeclContext moduleDecl() {
			return getRuleContext(ModuleDeclContext.class,0);
		}
		public FunctionDeclContext functionDecl() {
			return getRuleContext(FunctionDeclContext.class,0);
		}
		public StructDeclContext structDecl() {
			return getRuleContext(StructDeclContext.class,0);
		}
		public EnumDeclContext enumDecl() {
			return getRuleContext(EnumDeclContext.class,0);
		}
		public InterfaceDeclContext interfaceDecl() {
			return getRuleContext(InterfaceDeclContext.class,0);
		}
		public ImplDeclContext implDecl() {
			return getRuleContext(ImplDeclContext.class,0);
		}
		public ActorDeclContext actorDecl() {
			return getRuleContext(ActorDeclContext.class,0);
		}
		public TypeAliasContext typeAlias() {
			return getRuleContext(TypeAliasContext.class,0);
		}
		public ImportDeclContext importDecl() {
			return getRuleContext(ImportDeclContext.class,0);
		}
		public ConstantDeclContext constantDecl() {
			return getRuleContext(ConstantDeclContext.class,0);
		}
		public ItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_item; }
	}

	public final ItemContext item() throws RecognitionException {
		ItemContext _localctx = new ItemContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_item);
		try {
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(146);
				moduleDecl();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(147);
				functionDecl();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(148);
				structDecl();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(149);
				enumDecl();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(150);
				interfaceDecl();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(151);
				implDecl();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(152);
				actorDecl();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(153);
				typeAlias();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(154);
				importDecl();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(155);
				constantDecl();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleDeclContext extends ParserRuleContext {
		public TerminalNode MODULE() { return getToken(UnifiedParser.MODULE, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<ItemContext> item() {
			return getRuleContexts(ItemContext.class);
		}
		public ItemContext item(int i) {
			return getRuleContext(ItemContext.class,i);
		}
		public ModuleDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleDecl; }
	}

	public final ModuleDeclContext moduleDecl() throws RecognitionException {
		ModuleDeclContext _localctx = new ModuleDeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_moduleDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			match(MODULE);
			setState(159);
			identifier();
			setState(160);
			match(LBRACE);
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 52)) & ~0x3f) == 0 && ((1L << (_la - 52)) & 35835L) != 0)) {
				{
				{
				setState(161);
				item();
				}
				}
				setState(166);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(167);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImportDeclContext extends ParserRuleContext {
		public TerminalNode IMPORT() { return getToken(UnifiedParser.IMPORT, 0); }
		public ImportPathContext importPath() {
			return getRuleContext(ImportPathContext.class,0);
		}
		public TerminalNode AS() { return getToken(UnifiedParser.AS, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public ImportListContext importList() {
			return getRuleContext(ImportListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public ImportDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importDecl; }
	}

	public final ImportDeclContext importDecl() throws RecognitionException {
		ImportDeclContext _localctx = new ImportDeclContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_importDecl);
		int _la;
		try {
			setState(186);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				match(IMPORT);
				setState(170);
				importPath();
				setState(173);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AS) {
					{
					setState(171);
					match(AS);
					setState(172);
					identifier();
					}
				}

				setState(176);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(175);
					match(SEMI);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(178);
				match(IMPORT);
				setState(179);
				importPath();
				setState(180);
				match(LBRACE);
				setState(181);
				importList();
				setState(182);
				match(RBRACE);
				setState(184);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(183);
					match(SEMI);
					}
				}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImportPathContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(UnifiedParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(UnifiedParser.DOT, i);
		}
		public ImportPathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importPath; }
	}

	public final ImportPathContext importPath() throws RecognitionException {
		ImportPathContext _localctx = new ImportPathContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_importPath);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			identifier();
			setState(193);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(189);
				match(DOT);
				setState(190);
				identifier();
				}
				}
				setState(195);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImportListContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public List<TerminalNode> AS() { return getTokens(UnifiedParser.AS); }
		public TerminalNode AS(int i) {
			return getToken(UnifiedParser.AS, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public TerminalNode STAR() { return getToken(UnifiedParser.STAR, 0); }
		public ImportListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importList; }
	}

	public final ImportListContext importList() throws RecognitionException {
		ImportListContext _localctx = new ImportListContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_importList);
		int _la;
		try {
			setState(213);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				identifier();
				setState(199);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AS) {
					{
					setState(197);
					match(AS);
					setState(198);
					identifier();
					}
				}

				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(201);
					match(COMMA);
					setState(202);
					identifier();
					setState(205);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==AS) {
						{
						setState(203);
						match(AS);
						setState(204);
						identifier();
						}
					}

					}
					}
					setState(211);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case STAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(212);
				match(STAR);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionDeclContext extends ParserRuleContext {
		public TerminalNode FN() { return getToken(UnifiedParser.FN, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(UnifiedParser.ARROW, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public WhereClauseContext whereClause() {
			return getRuleContext(WhereClauseContext.class,0);
		}
		public FunctionDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDecl; }
	}

	public final FunctionDeclContext functionDecl() throws RecognitionException {
		FunctionDeclContext _localctx = new FunctionDeclContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_functionDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PUB) {
				{
				setState(215);
				match(PUB);
				}
			}

			setState(218);
			match(FN);
			setState(219);
			identifier();
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(220);
				genericParams();
				}
			}

			setState(223);
			match(LPAREN);
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SELF || _la==Identifier) {
				{
				setState(224);
				paramList();
				}
			}

			setState(227);
			match(RPAREN);
			setState(230);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARROW) {
				{
				setState(228);
				match(ARROW);
				setState(229);
				type(0);
				}
			}

			setState(233);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHERE) {
				{
				setState(232);
				whereClause();
				}
			}

			setState(235);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamListContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public ParamListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramList; }
	}

	public final ParamListContext paramList() throws RecognitionException {
		ParamListContext _localctx = new ParamListContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_paramList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			parameter();
			setState(242);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(238);
				match(COMMA);
				setState(239);
				parameter();
				}
				}
				setState(244);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public List<TerminalNode> SELF() { return getTokens(UnifiedParser.SELF); }
		public TerminalNode SELF(int i) {
			return getToken(UnifiedParser.SELF, i);
		}
		public TerminalNode BIT_AND() { return getToken(UnifiedParser.BIT_AND, 0); }
		public TerminalNode MUT() { return getToken(UnifiedParser.MUT, 0); }
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_parameter);
		try {
			setState(259);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(245);
				identifier();
				setState(246);
				match(COLON);
				setState(247);
				type(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(249);
				match(SELF);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(250);
				match(SELF);
				setState(251);
				match(COLON);
				setState(252);
				match(BIT_AND);
				setState(253);
				match(SELF);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(254);
				match(SELF);
				setState(255);
				match(COLON);
				setState(256);
				match(BIT_AND);
				setState(257);
				match(MUT);
				setState(258);
				match(SELF);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GenericParamsContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(UnifiedParser.LT, 0); }
		public List<GenericParamContext> genericParam() {
			return getRuleContexts(GenericParamContext.class);
		}
		public GenericParamContext genericParam(int i) {
			return getRuleContext(GenericParamContext.class,i);
		}
		public TerminalNode GT() { return getToken(UnifiedParser.GT, 0); }
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public GenericParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericParams; }
	}

	public final GenericParamsContext genericParams() throws RecognitionException {
		GenericParamsContext _localctx = new GenericParamsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_genericParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			match(LT);
			setState(262);
			genericParam();
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(263);
				match(COMMA);
				setState(264);
				genericParam();
				}
				}
				setState(269);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(270);
			match(GT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GenericParamContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeConstraintContext typeConstraint() {
			return getRuleContext(TypeConstraintContext.class,0);
		}
		public GenericParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_genericParam; }
	}

	public final GenericParamContext genericParam() throws RecognitionException {
		GenericParamContext _localctx = new GenericParamContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_genericParam);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			identifier();
			setState(275);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(273);
				match(COLON);
				setState(274);
				typeConstraint();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeConstraintContext extends ParserRuleContext {
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public List<TerminalNode> PLUS() { return getTokens(UnifiedParser.PLUS); }
		public TerminalNode PLUS(int i) {
			return getToken(UnifiedParser.PLUS, i);
		}
		public TypeConstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeConstraint; }
	}

	public final TypeConstraintContext typeConstraint() throws RecognitionException {
		TypeConstraintContext _localctx = new TypeConstraintContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_typeConstraint);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			type(0);
			setState(282);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS) {
				{
				{
				setState(278);
				match(PLUS);
				setState(279);
				type(0);
				}
				}
				setState(284);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhereClauseContext extends ParserRuleContext {
		public TerminalNode WHERE() { return getToken(UnifiedParser.WHERE, 0); }
		public List<WhereConstraintContext> whereConstraint() {
			return getRuleContexts(WhereConstraintContext.class);
		}
		public WhereConstraintContext whereConstraint(int i) {
			return getRuleContext(WhereConstraintContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public WhereClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whereClause; }
	}

	public final WhereClauseContext whereClause() throws RecognitionException {
		WhereClauseContext _localctx = new WhereClauseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_whereClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			match(WHERE);
			setState(286);
			whereConstraint();
			setState(291);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(287);
				match(COMMA);
				setState(288);
				whereConstraint();
				}
				}
				setState(293);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhereConstraintContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeConstraintContext typeConstraint() {
			return getRuleContext(TypeConstraintContext.class,0);
		}
		public WhereConstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whereConstraint; }
	}

	public final WhereConstraintContext whereConstraint() throws RecognitionException {
		WhereConstraintContext _localctx = new WhereConstraintContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_whereConstraint);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(294);
			type(0);
			setState(295);
			match(COLON);
			setState(296);
			typeConstraint();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclContext extends ParserRuleContext {
		public TerminalNode STRUCT() { return getToken(UnifiedParser.STRUCT, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public List<StructMemberContext> structMember() {
			return getRuleContexts(StructMemberContext.class);
		}
		public StructMemberContext structMember(int i) {
			return getRuleContext(StructMemberContext.class,i);
		}
		public StructDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDecl; }
	}

	public final StructDeclContext structDecl() throws RecognitionException {
		StructDeclContext _localctx = new StructDeclContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_structDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PUB) {
				{
				setState(298);
				match(PUB);
				}
			}

			setState(301);
			match(STRUCT);
			setState(302);
			identifier();
			setState(304);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(303);
				genericParams();
				}
			}

			setState(306);
			match(LBRACE);
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 55)) & ~0x3f) == 0 && ((1L << (_la - 55)) & 8589938689L) != 0)) {
				{
				{
				setState(307);
				structMember();
				}
				}
				setState(312);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(313);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructMemberContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public FunctionDeclContext functionDecl() {
			return getRuleContext(FunctionDeclContext.class,0);
		}
		public StructMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMember; }
	}

	public final StructMemberContext structMember() throws RecognitionException {
		StructMemberContext _localctx = new StructMemberContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_structMember);
		int _la;
		try {
			setState(323);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(316);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUB) {
					{
					setState(315);
					match(PUB);
					}
				}

				setState(318);
				identifier();
				setState(319);
				match(COLON);
				setState(320);
				type(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(322);
				functionDecl();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumDeclContext extends ParserRuleContext {
		public TerminalNode ENUM() { return getToken(UnifiedParser.ENUM, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public List<EnumVariantContext> enumVariant() {
			return getRuleContexts(EnumVariantContext.class);
		}
		public EnumVariantContext enumVariant(int i) {
			return getRuleContext(EnumVariantContext.class,i);
		}
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public EnumDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumDecl; }
	}

	public final EnumDeclContext enumDecl() throws RecognitionException {
		EnumDeclContext _localctx = new EnumDeclContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_enumDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PUB) {
				{
				setState(325);
				match(PUB);
				}
			}

			setState(328);
			match(ENUM);
			setState(329);
			identifier();
			setState(331);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(330);
				genericParams();
				}
			}

			setState(333);
			match(LBRACE);
			setState(334);
			enumVariant();
			setState(339);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(335);
				match(COMMA);
				setState(336);
				enumVariant();
				}
				}
				setState(341);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(342);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumVariantContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public EnumVariantParamsContext enumVariantParams() {
			return getRuleContext(EnumVariantParamsContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public EnumVariantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumVariant; }
	}

	public final EnumVariantContext enumVariant() throws RecognitionException {
		EnumVariantContext _localctx = new EnumVariantContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_enumVariant);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			identifier();
			setState(349);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(345);
				match(LPAREN);
				setState(346);
				enumVariantParams();
				setState(347);
				match(RPAREN);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumVariantParamsContext extends ParserRuleContext {
		public List<EnumVariantParamContext> enumVariantParam() {
			return getRuleContexts(EnumVariantParamContext.class);
		}
		public EnumVariantParamContext enumVariantParam(int i) {
			return getRuleContext(EnumVariantParamContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public EnumVariantParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumVariantParams; }
	}

	public final EnumVariantParamsContext enumVariantParams() throws RecognitionException {
		EnumVariantParamsContext _localctx = new EnumVariantParamsContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_enumVariantParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(351);
			enumVariantParam();
			setState(356);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(352);
				match(COMMA);
				setState(353);
				enumVariantParam();
				}
				}
				setState(358);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumVariantParamContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public EnumVariantParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumVariantParam; }
	}

	public final EnumVariantParamContext enumVariantParam() throws RecognitionException {
		EnumVariantParamContext _localctx = new EnumVariantParamContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_enumVariantParam);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(362);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(359);
				identifier();
				setState(360);
				match(COLON);
				}
				break;
			}
			setState(364);
			type(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceDeclContext extends ParserRuleContext {
		public TerminalNode INTERFACE() { return getToken(UnifiedParser.INTERFACE, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public List<InterfaceMemberContext> interfaceMember() {
			return getRuleContexts(InterfaceMemberContext.class);
		}
		public InterfaceMemberContext interfaceMember(int i) {
			return getRuleContext(InterfaceMemberContext.class,i);
		}
		public InterfaceDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceDecl; }
	}

	public final InterfaceDeclContext interfaceDecl() throws RecognitionException {
		InterfaceDeclContext _localctx = new InterfaceDeclContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_interfaceDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(367);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PUB) {
				{
				setState(366);
				match(PUB);
				}
			}

			setState(369);
			match(INTERFACE);
			setState(370);
			identifier();
			setState(372);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(371);
				genericParams();
				}
			}

			setState(374);
			match(LBRACE);
			setState(378);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FN || _la==TYPE) {
				{
				{
				setState(375);
				interfaceMember();
				}
				}
				setState(380);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(381);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceMemberContext extends ParserRuleContext {
		public FunctionSigContext functionSig() {
			return getRuleContext(FunctionSigContext.class,0);
		}
		public TerminalNode TYPE() { return getToken(UnifiedParser.TYPE, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public InterfaceMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceMember; }
	}

	public final InterfaceMemberContext interfaceMember() throws RecognitionException {
		InterfaceMemberContext _localctx = new InterfaceMemberContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_interfaceMember);
		int _la;
		try {
			setState(389);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FN:
				enterOuterAlt(_localctx, 1);
				{
				setState(383);
				functionSig();
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(384);
				match(TYPE);
				setState(385);
				identifier();
				setState(387);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(386);
					match(SEMI);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionSigContext extends ParserRuleContext {
		public TerminalNode FN() { return getToken(UnifiedParser.FN, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(UnifiedParser.ARROW, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FunctionSigContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionSig; }
	}

	public final FunctionSigContext functionSig() throws RecognitionException {
		FunctionSigContext _localctx = new FunctionSigContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_functionSig);
		int _la;
		try {
			setState(424);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,46,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(391);
				match(FN);
				setState(392);
				identifier();
				setState(394);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(393);
					genericParams();
					}
				}

				setState(396);
				match(LPAREN);
				setState(398);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SELF || _la==Identifier) {
					{
					setState(397);
					paramList();
					}
				}

				setState(400);
				match(RPAREN);
				setState(403);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ARROW) {
					{
					setState(401);
					match(ARROW);
					setState(402);
					type(0);
					}
				}

				setState(406);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(405);
					match(SEMI);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(408);
				match(FN);
				setState(409);
				identifier();
				setState(411);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(410);
					genericParams();
					}
				}

				setState(413);
				match(LPAREN);
				setState(415);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SELF || _la==Identifier) {
					{
					setState(414);
					paramList();
					}
				}

				setState(417);
				match(RPAREN);
				setState(420);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ARROW) {
					{
					setState(418);
					match(ARROW);
					setState(419);
					type(0);
					}
				}

				setState(422);
				block();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImplDeclContext extends ParserRuleContext {
		public TerminalNode IMPL() { return getToken(UnifiedParser.IMPL, 0); }
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public TerminalNode FOR() { return getToken(UnifiedParser.FOR, 0); }
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public WhereClauseContext whereClause() {
			return getRuleContext(WhereClauseContext.class,0);
		}
		public List<ImplMemberContext> implMember() {
			return getRuleContexts(ImplMemberContext.class);
		}
		public ImplMemberContext implMember(int i) {
			return getRuleContext(ImplMemberContext.class,i);
		}
		public ImplDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implDecl; }
	}

	public final ImplDeclContext implDecl() throws RecognitionException {
		ImplDeclContext _localctx = new ImplDeclContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_implDecl);
		int _la;
		try {
			setState(462);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,53,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(426);
				match(IMPL);
				setState(428);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(427);
					genericParams();
					}
				}

				setState(430);
				type(0);
				setState(431);
				match(FOR);
				setState(432);
				type(0);
				setState(434);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHERE) {
					{
					setState(433);
					whereClause();
					}
				}

				setState(436);
				match(LBRACE);
				setState(440);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 55)) & ~0x3f) == 0 && ((1L << (_la - 55)) & 4161L) != 0)) {
					{
					{
					setState(437);
					implMember();
					}
					}
					setState(442);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(443);
				match(RBRACE);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(445);
				match(IMPL);
				setState(447);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(446);
					genericParams();
					}
				}

				setState(449);
				type(0);
				setState(451);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHERE) {
					{
					setState(450);
					whereClause();
					}
				}

				setState(453);
				match(LBRACE);
				setState(457);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 55)) & ~0x3f) == 0 && ((1L << (_la - 55)) & 4161L) != 0)) {
					{
					{
					setState(454);
					implMember();
					}
					}
					setState(459);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(460);
				match(RBRACE);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImplMemberContext extends ParserRuleContext {
		public FunctionDeclContext functionDecl() {
			return getRuleContext(FunctionDeclContext.class,0);
		}
		public TerminalNode TYPE() { return getToken(UnifiedParser.TYPE, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public ImplMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implMember; }
	}

	public final ImplMemberContext implMember() throws RecognitionException {
		ImplMemberContext _localctx = new ImplMemberContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_implMember);
		int _la;
		try {
			setState(472);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FN:
			case PUB:
				enterOuterAlt(_localctx, 1);
				{
				setState(464);
				functionDecl();
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(465);
				match(TYPE);
				setState(466);
				identifier();
				setState(467);
				match(ASSIGN);
				setState(468);
				type(0);
				setState(470);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(469);
					match(SEMI);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ActorDeclContext extends ParserRuleContext {
		public TerminalNode ACTOR() { return getToken(UnifiedParser.ACTOR, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public List<ActorMemberContext> actorMember() {
			return getRuleContexts(ActorMemberContext.class);
		}
		public ActorMemberContext actorMember(int i) {
			return getRuleContext(ActorMemberContext.class,i);
		}
		public ActorDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_actorDecl; }
	}

	public final ActorDeclContext actorDecl() throws RecognitionException {
		ActorDeclContext _localctx = new ActorDeclContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_actorDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(475);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PUB) {
				{
				setState(474);
				match(PUB);
				}
			}

			setState(477);
			match(ACTOR);
			setState(478);
			identifier();
			setState(480);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(479);
				genericParams();
				}
			}

			setState(482);
			match(LBRACE);
			setState(486);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 55)) & ~0x3f) == 0 && ((1L << (_la - 55)) & 5121L) != 0)) {
				{
				{
				setState(483);
				actorMember();
				}
				}
				setState(488);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(489);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ActorMemberContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(UnifiedParser.VAR, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public FunctionDeclContext functionDecl() {
			return getRuleContext(FunctionDeclContext.class,0);
		}
		public ActorMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_actorMember; }
	}

	public final ActorMemberContext actorMember() throws RecognitionException {
		ActorMemberContext _localctx = new ActorMemberContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_actorMember);
		int _la;
		try {
			setState(506);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,62,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(492);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUB) {
					{
					setState(491);
					match(PUB);
					}
				}

				setState(494);
				match(VAR);
				setState(495);
				identifier();
				setState(496);
				match(COLON);
				setState(497);
				type(0);
				setState(500);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(498);
					match(ASSIGN);
					setState(499);
					expr(0);
					}
				}

				setState(503);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(502);
					match(SEMI);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(505);
				functionDecl();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeAliasContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(UnifiedParser.TYPE, 0); }
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public GenericParamsContext genericParams() {
			return getRuleContext(GenericParamsContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public TerminalNode REFINE() { return getToken(UnifiedParser.REFINE, 0); }
		public List<TerminalNode> BIT_OR() { return getTokens(UnifiedParser.BIT_OR); }
		public TerminalNode BIT_OR(int i) {
			return getToken(UnifiedParser.BIT_OR, i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TypeAliasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAlias; }
	}

	public final TypeAliasContext typeAlias() throws RecognitionException {
		TypeAliasContext _localctx = new TypeAliasContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_typeAlias);
		int _la;
		try {
			setState(539);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,69,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(509);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUB) {
					{
					setState(508);
					match(PUB);
					}
				}

				setState(511);
				match(TYPE);
				setState(512);
				identifier();
				setState(514);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(513);
					genericParams();
					}
				}

				setState(516);
				match(ASSIGN);
				setState(517);
				type(0);
				setState(519);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(518);
					match(SEMI);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(522);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUB) {
					{
					setState(521);
					match(PUB);
					}
				}

				setState(524);
				match(TYPE);
				setState(525);
				identifier();
				setState(527);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LT) {
					{
					setState(526);
					genericParams();
					}
				}

				setState(529);
				match(ASSIGN);
				setState(530);
				type(0);
				setState(531);
				match(REFINE);
				setState(532);
				match(BIT_OR);
				setState(533);
				identifier();
				setState(534);
				match(BIT_OR);
				setState(535);
				expr(0);
				setState(537);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMI) {
					{
					setState(536);
					match(SEMI);
					}
				}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConstantDeclContext extends ParserRuleContext {
		public TerminalNode CONST() { return getToken(UnifiedParser.CONST, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PUB() { return getToken(UnifiedParser.PUB, 0); }
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public ConstantDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantDecl; }
	}

	public final ConstantDeclContext constantDecl() throws RecognitionException {
		ConstantDeclContext _localctx = new ConstantDeclContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_constantDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(542);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PUB) {
				{
				setState(541);
				match(PUB);
				}
			}

			setState(544);
			match(CONST);
			setState(545);
			identifier();
			setState(546);
			match(COLON);
			setState(547);
			type(0);
			setState(548);
			match(ASSIGN);
			setState(549);
			expr(0);
			setState(551);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(550);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public TerminalNode LT() { return getToken(UnifiedParser.LT, 0); }
		public TypeListContext typeList() {
			return getRuleContext(TypeListContext.class,0);
		}
		public TerminalNode GT() { return getToken(UnifiedParser.GT, 0); }
		public TerminalNode FN() { return getToken(UnifiedParser.FN, 0); }
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public TerminalNode ARROW() { return getToken(UnifiedParser.ARROW, 0); }
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public TerminalNode BIT_AND() { return getToken(UnifiedParser.BIT_AND, 0); }
		public TerminalNode MUT() { return getToken(UnifiedParser.MUT, 0); }
		public TerminalNode IMPL() { return getToken(UnifiedParser.IMPL, 0); }
		public TerminalNode BIT_OR() { return getToken(UnifiedParser.BIT_OR, 0); }
		public List<TerminalNode> DOUBLE_COLON() { return getTokens(UnifiedParser.DOUBLE_COLON); }
		public TerminalNode DOUBLE_COLON(int i) {
			return getToken(UnifiedParser.DOUBLE_COLON, i);
		}
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		return type(0);
	}

	private TypeContext type(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		TypeContext _localctx = new TypeContext(_ctx, _parentState);
		TypeContext _prevctx = _localctx;
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_type, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(580);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,74,_ctx) ) {
			case 1:
				{
				setState(554);
				identifier();
				setState(559);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,72,_ctx) ) {
				case 1:
					{
					setState(555);
					match(LT);
					setState(556);
					typeList();
					setState(557);
					match(GT);
					}
					break;
				}
				}
				break;
			case 2:
				{
				setState(561);
				match(FN);
				setState(562);
				match(LPAREN);
				setState(564);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 612489549322649608L) != 0) || _la==Identifier) {
					{
					setState(563);
					typeList();
					}
				}

				setState(566);
				match(RPAREN);
				setState(567);
				match(ARROW);
				setState(568);
				type(7);
				}
				break;
			case 3:
				{
				setState(569);
				match(LPAREN);
				setState(570);
				typeList();
				setState(571);
				match(RPAREN);
				}
				break;
			case 4:
				{
				setState(573);
				match(BIT_AND);
				setState(574);
				type(5);
				}
				break;
			case 5:
				{
				setState(575);
				match(BIT_AND);
				setState(576);
				match(MUT);
				setState(577);
				type(4);
				}
				break;
			case 6:
				{
				setState(578);
				match(IMPL);
				setState(579);
				identifier();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(594);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,77,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(592);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,76,_ctx) ) {
					case 1:
						{
						_localctx = new TypeContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type);
						setState(582);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(583);
						match(BIT_OR);
						setState(584);
						type(4);
						}
						break;
					case 2:
						{
						_localctx = new TypeContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_type);
						setState(585);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(588); 
						_errHandler.sync(this);
						_alt = 1;
						do {
							switch (_alt) {
							case 1:
								{
								{
								setState(586);
								match(DOUBLE_COLON);
								setState(587);
								identifier();
								}
								}
								break;
							default:
								throw new NoViableAltException(this);
							}
							setState(590); 
							_errHandler.sync(this);
							_alt = getInterpreter().adaptivePredict(_input,75,_ctx);
						} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
						}
						break;
					}
					} 
				}
				setState(596);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,77,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeListContext extends ParserRuleContext {
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public TypeListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeList; }
	}

	public final TypeListContext typeList() throws RecognitionException {
		TypeListContext _localctx = new TypeListContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_typeList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(597);
			type(0);
			setState(602);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(598);
				match(COMMA);
				setState(599);
				type(0);
				}
				}
				setState(604);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public LetStatementContext letStatement() {
			return getRuleContext(LetStatementContext.class,0);
		}
		public VarStatementContext varStatement() {
			return getRuleContext(VarStatementContext.class,0);
		}
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public RegionStatementContext regionStatement() {
			return getRuleContext(RegionStatementContext.class,0);
		}
		public ExprStatementContext exprStatement() {
			return getRuleContext(ExprStatementContext.class,0);
		}
		public ReturnStatementContext returnStatement() {
			return getRuleContext(ReturnStatementContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public LoopStatementContext loopStatement() {
			return getRuleContext(LoopStatementContext.class,0);
		}
		public WhileStatementContext whileStatement() {
			return getRuleContext(WhileStatementContext.class,0);
		}
		public ForStatementContext forStatement() {
			return getRuleContext(ForStatementContext.class,0);
		}
		public SwitchStatementContext switchStatement() {
			return getRuleContext(SwitchStatementContext.class,0);
		}
		public BreakStatementContext breakStatement() {
			return getRuleContext(BreakStatementContext.class,0);
		}
		public ContinueStatementContext continueStatement() {
			return getRuleContext(ContinueStatementContext.class,0);
		}
		public BlockStatementContext blockStatement() {
			return getRuleContext(BlockStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_statement);
		try {
			setState(619);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,79,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(605);
				letStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(606);
				varStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(607);
				assignmentStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(608);
				regionStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(609);
				exprStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(610);
				returnStatement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(611);
				ifStatement();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(612);
				loopStatement();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(613);
				whileStatement();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(614);
				forStatement();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(615);
				switchStatement();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(616);
				breakStatement();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(617);
				continueStatement();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(618);
				blockStatement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LetStatementContext extends ParserRuleContext {
		public TerminalNode LET() { return getToken(UnifiedParser.LET, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode MUT() { return getToken(UnifiedParser.MUT, 0); }
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public LetStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_letStatement; }
	}

	public final LetStatementContext letStatement() throws RecognitionException {
		LetStatementContext _localctx = new LetStatementContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_letStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(621);
			match(LET);
			setState(623);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==MUT) {
				{
				setState(622);
				match(MUT);
				}
			}

			setState(625);
			identifier();
			setState(628);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(626);
				match(COLON);
				setState(627);
				type(0);
				}
			}

			setState(630);
			match(ASSIGN);
			setState(631);
			expr(0);
			setState(633);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(632);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarStatementContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(UnifiedParser.VAR, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public VarStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varStatement; }
	}

	public final VarStatementContext varStatement() throws RecognitionException {
		VarStatementContext _localctx = new VarStatementContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_varStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(635);
			match(VAR);
			setState(636);
			identifier();
			setState(637);
			match(COLON);
			setState(638);
			type(0);
			setState(641);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(639);
				match(ASSIGN);
				setState(640);
				expr(0);
				}
			}

			setState(644);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(643);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public AssignmentOpContext assignmentOp() {
			return getRuleContext(AssignmentOpContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_assignmentStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(646);
			identifier();
			setState(647);
			assignmentOp();
			setState(648);
			expr(0);
			setState(650);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(649);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentOpContext extends ParserRuleContext {
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public TerminalNode PLUS_ASSIGN() { return getToken(UnifiedParser.PLUS_ASSIGN, 0); }
		public TerminalNode MINUS_ASSIGN() { return getToken(UnifiedParser.MINUS_ASSIGN, 0); }
		public TerminalNode STAR_ASSIGN() { return getToken(UnifiedParser.STAR_ASSIGN, 0); }
		public TerminalNode DIV_ASSIGN() { return getToken(UnifiedParser.DIV_ASSIGN, 0); }
		public TerminalNode MOD_ASSIGN() { return getToken(UnifiedParser.MOD_ASSIGN, 0); }
		public AssignmentOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentOp; }
	}

	public final AssignmentOpContext assignmentOp() throws RecognitionException {
		AssignmentOpContext _localctx = new AssignmentOpContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_assignmentOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(652);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 541165879296L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RegionStatementContext extends ParserRuleContext {
		public TerminalNode REGION() { return getToken(UnifiedParser.REGION, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public RegionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_regionStatement; }
	}

	public final RegionStatementContext regionStatement() throws RecognitionException {
		RegionStatementContext _localctx = new RegionStatementContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_regionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(654);
			match(REGION);
			setState(655);
			identifier();
			setState(656);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprStatementContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public ExprStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprStatement; }
	}

	public final ExprStatementContext exprStatement() throws RecognitionException {
		ExprStatementContext _localctx = new ExprStatementContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_exprStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(658);
			expr(0);
			setState(660);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(659);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(UnifiedParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public ReturnStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStatement; }
	}

	public final ReturnStatementContext returnStatement() throws RecognitionException {
		ReturnStatementContext _localctx = new ReturnStatementContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_returnStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(662);
			match(RETURN);
			setState(664);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,87,_ctx) ) {
			case 1:
				{
				setState(663);
				expr(0);
				}
				break;
			}
			setState(667);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(666);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public List<TerminalNode> IF() { return getTokens(UnifiedParser.IF); }
		public TerminalNode IF(int i) {
			return getToken(UnifiedParser.IF, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> ELSE() { return getTokens(UnifiedParser.ELSE); }
		public TerminalNode ELSE(int i) {
			return getToken(UnifiedParser.ELSE, i);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_ifStatement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(669);
			match(IF);
			setState(670);
			expr(0);
			setState(671);
			block();
			setState(679);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,89,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(672);
					match(ELSE);
					setState(673);
					match(IF);
					setState(674);
					expr(0);
					setState(675);
					block();
					}
					} 
				}
				setState(681);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,89,_ctx);
			}
			setState(684);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(682);
				match(ELSE);
				setState(683);
				block();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LoopStatementContext extends ParserRuleContext {
		public TerminalNode LOOP() { return getToken(UnifiedParser.LOOP, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public LoopStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loopStatement; }
	}

	public final LoopStatementContext loopStatement() throws RecognitionException {
		LoopStatementContext _localctx = new LoopStatementContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_loopStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(689);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(686);
				identifier();
				setState(687);
				match(COLON);
				}
			}

			setState(691);
			match(LOOP);
			setState(692);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhileStatementContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(UnifiedParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public WhileStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileStatement; }
	}

	public final WhileStatementContext whileStatement() throws RecognitionException {
		WhileStatementContext _localctx = new WhileStatementContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_whileStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(697);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(694);
				identifier();
				setState(695);
				match(COLON);
				}
			}

			setState(699);
			match(WHILE);
			setState(700);
			expr(0);
			setState(701);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForStatementContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(UnifiedParser.FOR, 0); }
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public TerminalNode IN() { return getToken(UnifiedParser.IN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public ForStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStatement; }
	}

	public final ForStatementContext forStatement() throws RecognitionException {
		ForStatementContext _localctx = new ForStatementContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_forStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(706);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(703);
				identifier();
				setState(704);
				match(COLON);
				}
			}

			setState(708);
			match(FOR);
			setState(709);
			identifier();
			setState(710);
			match(IN);
			setState(711);
			expr(0);
			setState(712);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStatementContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(UnifiedParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<CaseClauseContext> caseClause() {
			return getRuleContexts(CaseClauseContext.class);
		}
		public CaseClauseContext caseClause(int i) {
			return getRuleContext(CaseClauseContext.class,i);
		}
		public SwitchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStatement; }
	}

	public final SwitchStatementContext switchStatement() throws RecognitionException {
		SwitchStatementContext _localctx = new SwitchStatementContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_switchStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(714);
			match(SWITCH);
			setState(715);
			expr(0);
			setState(716);
			match(LBRACE);
			setState(720);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(717);
				caseClause();
				}
				}
				setState(722);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(723);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CaseClauseContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(UnifiedParser.CASE, 0); }
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(UnifiedParser.ARROW, 0); }
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public CaseClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseClause; }
	}

	public final CaseClauseContext caseClause() throws RecognitionException {
		CaseClauseContext _localctx = new CaseClauseContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_caseClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(725);
			match(CASE);
			setState(726);
			pattern(0);
			setState(727);
			_la = _input.LA(1);
			if ( !(_la==COLON || _la==ARROW) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(730);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,95,_ctx) ) {
			case 1:
				{
				setState(728);
				statement();
				}
				break;
			case 2:
				{
				setState(729);
				block();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(UnifiedParser.BREAK, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public BreakStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakStatement; }
	}

	public final BreakStatementContext breakStatement() throws RecognitionException {
		BreakStatementContext _localctx = new BreakStatementContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_breakStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(732);
			match(BREAK);
			setState(734);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,96,_ctx) ) {
			case 1:
				{
				setState(733);
				identifier();
				}
				break;
			}
			setState(737);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(736);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(UnifiedParser.CONTINUE, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(UnifiedParser.SEMI, 0); }
		public ContinueStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continueStatement; }
	}

	public final ContinueStatementContext continueStatement() throws RecognitionException {
		ContinueStatementContext _localctx = new ContinueStatementContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_continueStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(739);
			match(CONTINUE);
			setState(741);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,98,_ctx) ) {
			case 1:
				{
				setState(740);
				identifier();
				}
				break;
			}
			setState(744);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(743);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockStatementContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public BlockStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockStatement; }
	}

	public final BlockStatementContext blockStatement() throws RecognitionException {
		BlockStatementContext _localctx = new BlockStatementContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_blockStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(746);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PatternContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode UNDERSCORE() { return getToken(UnifiedParser.UNDERSCORE, 0); }
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public PatternListContext patternList() {
			return getRuleContext(PatternListContext.class,0);
		}
		public TerminalNode LET() { return getToken(UnifiedParser.LET, 0); }
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode IF() { return getToken(UnifiedParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public FieldPatternListContext fieldPatternList() {
			return getRuleContext(FieldPatternListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<PatternContext> pattern() {
			return getRuleContexts(PatternContext.class);
		}
		public PatternContext pattern(int i) {
			return getRuleContext(PatternContext.class,i);
		}
		public TerminalNode RANGE() { return getToken(UnifiedParser.RANGE, 0); }
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
	}

	public final PatternContext pattern() throws RecognitionException {
		return pattern(0);
	}

	private PatternContext pattern(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PatternContext _localctx = new PatternContext(_ctx, _parentState);
		PatternContext _prevctx = _localctx;
		int _startState = 96;
		enterRecursionRule(_localctx, 96, RULE_pattern, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(779);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,104,_ctx) ) {
			case 1:
				{
				setState(749);
				identifier();
				}
				break;
			case 2:
				{
				setState(750);
				match(UNDERSCORE);
				}
				break;
			case 3:
				{
				setState(751);
				literal();
				}
				break;
			case 4:
				{
				setState(752);
				identifier();
				setState(753);
				match(LPAREN);
				setState(755);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LPAREN || _la==UNDERSCORE || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 532676609L) != 0)) {
					{
					setState(754);
					patternList();
					}
				}

				setState(757);
				match(RPAREN);
				}
				break;
			case 5:
				{
				setState(759);
				match(LET);
				setState(760);
				identifier();
				setState(763);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,101,_ctx) ) {
				case 1:
					{
					setState(761);
					match(COLON);
					setState(762);
					type(0);
					}
					break;
				}
				setState(767);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,102,_ctx) ) {
				case 1:
					{
					setState(765);
					match(IF);
					setState(766);
					expr(0);
					}
					break;
				}
				}
				break;
			case 6:
				{
				setState(769);
				identifier();
				setState(770);
				match(LBRACE);
				setState(771);
				fieldPatternList();
				setState(772);
				match(RBRACE);
				}
				break;
			case 7:
				{
				setState(774);
				match(LPAREN);
				setState(776);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LPAREN || _la==UNDERSCORE || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 532676609L) != 0)) {
					{
					setState(775);
					patternList();
					}
				}

				setState(778);
				match(RPAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(786);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,105,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PatternContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pattern);
					setState(781);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(782);
					match(RANGE);
					setState(783);
					pattern(6);
					}
					} 
				}
				setState(788);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,105,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PatternListContext extends ParserRuleContext {
		public List<PatternContext> pattern() {
			return getRuleContexts(PatternContext.class);
		}
		public PatternContext pattern(int i) {
			return getRuleContext(PatternContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public PatternListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_patternList; }
	}

	public final PatternListContext patternList() throws RecognitionException {
		PatternListContext _localctx = new PatternListContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_patternList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(789);
			pattern(0);
			setState(794);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(790);
				match(COMMA);
				setState(791);
				pattern(0);
				}
				}
				setState(796);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldPatternContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode RANGE() { return getToken(UnifiedParser.RANGE, 0); }
		public FieldPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldPattern; }
	}

	public final FieldPatternContext fieldPattern() throws RecognitionException {
		FieldPatternContext _localctx = new FieldPatternContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_fieldPattern);
		try {
			setState(803);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,107,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(797);
				identifier();
				setState(798);
				match(COLON);
				setState(799);
				pattern(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(801);
				identifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(802);
				match(RANGE);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldPatternListContext extends ParserRuleContext {
		public List<FieldPatternContext> fieldPattern() {
			return getRuleContexts(FieldPatternContext.class);
		}
		public FieldPatternContext fieldPattern(int i) {
			return getRuleContext(FieldPatternContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public FieldPatternListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldPatternList; }
	}

	public final FieldPatternListContext fieldPatternList() throws RecognitionException {
		FieldPatternListContext _localctx = new FieldPatternListContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_fieldPatternList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(805);
			fieldPattern();
			setState(810);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(806);
				match(COMMA);
				setState(807);
				fieldPattern();
				}
				}
				setState(812);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public PrimaryContext primary() {
			return getRuleContext(PrimaryContext.class,0);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode INC() { return getToken(UnifiedParser.INC, 0); }
		public TerminalNode DEC() { return getToken(UnifiedParser.DEC, 0); }
		public TerminalNode PLUS() { return getToken(UnifiedParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(UnifiedParser.MINUS, 0); }
		public TerminalNode NOT() { return getToken(UnifiedParser.NOT, 0); }
		public TerminalNode BIT_NOT() { return getToken(UnifiedParser.BIT_NOT, 0); }
		public TerminalNode MOVE() { return getToken(UnifiedParser.MOVE, 0); }
		public TerminalNode AWAIT() { return getToken(UnifiedParser.AWAIT, 0); }
		public LambdaExprContext lambdaExpr() {
			return getRuleContext(LambdaExprContext.class,0);
		}
		public AsyncExprContext asyncExpr() {
			return getRuleContext(AsyncExprContext.class,0);
		}
		public List<TerminalNode> IF() { return getTokens(UnifiedParser.IF); }
		public TerminalNode IF(int i) {
			return getToken(UnifiedParser.IF, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> ELSE() { return getTokens(UnifiedParser.ELSE); }
		public TerminalNode ELSE(int i) {
			return getToken(UnifiedParser.ELSE, i);
		}
		public TerminalNode SWITCH() { return getToken(UnifiedParser.SWITCH, 0); }
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<CaseExprContext> caseExpr() {
			return getRuleContexts(CaseExprContext.class);
		}
		public CaseExprContext caseExpr(int i) {
			return getRuleContext(CaseExprContext.class,i);
		}
		public TerminalNode STAR() { return getToken(UnifiedParser.STAR, 0); }
		public TerminalNode DIV() { return getToken(UnifiedParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(UnifiedParser.MOD, 0); }
		public TerminalNode LSHIFT() { return getToken(UnifiedParser.LSHIFT, 0); }
		public TerminalNode RSHIFT() { return getToken(UnifiedParser.RSHIFT, 0); }
		public TerminalNode LT() { return getToken(UnifiedParser.LT, 0); }
		public TerminalNode LE() { return getToken(UnifiedParser.LE, 0); }
		public TerminalNode GT() { return getToken(UnifiedParser.GT, 0); }
		public TerminalNode GE() { return getToken(UnifiedParser.GE, 0); }
		public TerminalNode EQ() { return getToken(UnifiedParser.EQ, 0); }
		public TerminalNode NE() { return getToken(UnifiedParser.NE, 0); }
		public TerminalNode BIT_AND() { return getToken(UnifiedParser.BIT_AND, 0); }
		public TerminalNode BIT_XOR() { return getToken(UnifiedParser.BIT_XOR, 0); }
		public TerminalNode BIT_OR() { return getToken(UnifiedParser.BIT_OR, 0); }
		public TerminalNode AND() { return getToken(UnifiedParser.AND, 0); }
		public TerminalNode OR() { return getToken(UnifiedParser.OR, 0); }
		public TerminalNode QUESTION_QUESTION() { return getToken(UnifiedParser.QUESTION_QUESTION, 0); }
		public TerminalNode QUESTION() { return getToken(UnifiedParser.QUESTION, 0); }
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public TerminalNode ASSIGN() { return getToken(UnifiedParser.ASSIGN, 0); }
		public TerminalNode PLUS_ASSIGN() { return getToken(UnifiedParser.PLUS_ASSIGN, 0); }
		public TerminalNode MINUS_ASSIGN() { return getToken(UnifiedParser.MINUS_ASSIGN, 0); }
		public TerminalNode STAR_ASSIGN() { return getToken(UnifiedParser.STAR_ASSIGN, 0); }
		public TerminalNode DIV_ASSIGN() { return getToken(UnifiedParser.DIV_ASSIGN, 0); }
		public TerminalNode MOD_ASSIGN() { return getToken(UnifiedParser.MOD_ASSIGN, 0); }
		public TerminalNode LSHIFT_ASSIGN() { return getToken(UnifiedParser.LSHIFT_ASSIGN, 0); }
		public TerminalNode RSHIFT_ASSIGN() { return getToken(UnifiedParser.RSHIFT_ASSIGN, 0); }
		public TerminalNode AND_ASSIGN() { return getToken(UnifiedParser.AND_ASSIGN, 0); }
		public TerminalNode XOR_ASSIGN() { return getToken(UnifiedParser.XOR_ASSIGN, 0); }
		public TerminalNode OR_ASSIGN() { return getToken(UnifiedParser.OR_ASSIGN, 0); }
		public TerminalNode RANGE() { return getToken(UnifiedParser.RANGE, 0); }
		public TerminalNode RANGE_INCL() { return getToken(UnifiedParser.RANGE_INCL, 0); }
		public TerminalNode NULL_COND() { return getToken(UnifiedParser.NULL_COND, 0); }
		public TerminalNode DOT() { return getToken(UnifiedParser.DOT, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode IntLiteral() { return getToken(UnifiedParser.IntLiteral, 0); }
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public ArgListContext argList() {
			return getRuleContext(ArgListContext.class,0);
		}
		public TerminalNode LBRACK() { return getToken(UnifiedParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(UnifiedParser.RBRACK, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 104;
		enterRecursionRule(_localctx, 104, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(851);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,112,_ctx) ) {
			case 1:
				{
				setState(814);
				primary();
				}
				break;
			case 2:
				{
				setState(815);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849929494528L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(816);
				expr(21);
				}
				break;
			case 3:
				{
				setState(817);
				match(MOVE);
				setState(818);
				expr(20);
				}
				break;
			case 4:
				{
				setState(819);
				match(AWAIT);
				setState(820);
				expr(19);
				}
				break;
			case 5:
				{
				setState(821);
				lambdaExpr();
				}
				break;
			case 6:
				{
				setState(822);
				asyncExpr();
				}
				break;
			case 7:
				{
				setState(823);
				match(IF);
				setState(824);
				expr(0);
				setState(825);
				block();
				setState(833);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,109,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(826);
						match(ELSE);
						setState(827);
						match(IF);
						setState(828);
						expr(0);
						setState(829);
						block();
						}
						} 
					}
					setState(835);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,109,_ctx);
				}
				setState(838);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,110,_ctx) ) {
				case 1:
					{
					setState(836);
					match(ELSE);
					setState(837);
					block();
					}
					break;
				}
				}
				break;
			case 8:
				{
				setState(840);
				match(SWITCH);
				setState(841);
				expr(0);
				setState(842);
				match(LBRACE);
				setState(846);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==CASE) {
					{
					{
					setState(843);
					caseExpr();
					}
					}
					setState(848);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(849);
				match(RBRACE);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(927);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,117,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(925);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,116,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(853);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(854);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 229376L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(855);
						expr(19);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(856);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(857);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(858);
						expr(18);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(859);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(860);
						_la = _input.LA(1);
						if ( !(_la==LSHIFT || _la==RSHIFT) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(861);
						expr(17);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(862);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(863);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 8053063680L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(864);
						expr(16);
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(865);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(866);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NE) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(867);
						expr(15);
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(868);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(869);
						match(BIT_AND);
						setState(870);
						expr(14);
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(871);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(872);
						match(BIT_XOR);
						setState(873);
						expr(13);
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(874);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(875);
						match(BIT_OR);
						setState(876);
						expr(12);
						}
						break;
					case 9:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(877);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(878);
						match(AND);
						setState(879);
						expr(11);
						}
						break;
					case 10:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(880);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(881);
						match(OR);
						setState(882);
						expr(10);
						}
						break;
					case 11:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(883);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(884);
						match(QUESTION_QUESTION);
						setState(885);
						expr(8);
						}
						break;
					case 12:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(886);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(887);
						match(QUESTION);
						setState(888);
						expr(0);
						setState(889);
						match(COLON);
						setState(890);
						expr(7);
						}
						break;
					case 13:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(892);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(893);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 510164805353472L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(894);
						expr(6);
						}
						break;
					case 14:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(895);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(896);
						match(DOT);
						setState(899);
						_errHandler.sync(this);
						switch (_input.LA(1)) {
						case Identifier:
							{
							setState(897);
							identifier();
							}
							break;
						case IntLiteral:
							{
							setState(898);
							match(IntLiteral);
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						}
						break;
					case 15:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(901);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(902);
						match(DOT);
						setState(903);
						identifier();
						setState(904);
						match(LPAREN);
						setState(906);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
							{
							setState(905);
							argList();
							}
						}

						setState(908);
						match(RPAREN);
						}
						break;
					case 16:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(910);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(911);
						match(LBRACK);
						setState(912);
						expr(0);
						setState(913);
						match(RBRACK);
						}
						break;
					case 17:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(915);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(916);
						match(LPAREN);
						setState(918);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
							{
							setState(917);
							argList();
							}
						}

						setState(920);
						match(RPAREN);
						}
						break;
					case 18:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(921);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(922);
						_la = _input.LA(1);
						if ( !(_la==INC || _la==DEC) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						}
						break;
					case 19:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(923);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(924);
						match(QUESTION);
						}
						break;
					}
					} 
				}
				setState(929);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,117,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CaseExprContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(UnifiedParser.CASE, 0); }
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(UnifiedParser.ARROW, 0); }
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public CaseExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseExpr; }
	}

	public final CaseExprContext caseExpr() throws RecognitionException {
		CaseExprContext _localctx = new CaseExprContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_caseExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(930);
			match(CASE);
			setState(931);
			pattern(0);
			setState(932);
			_la = _input.LA(1);
			if ( !(_la==COLON || _la==ARROW) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(933);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode SELF() { return getToken(UnifiedParser.SELF, 0); }
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public StructExprContext structExpr() {
			return getRuleContext(StructExprContext.class,0);
		}
		public ListExprContext listExpr() {
			return getRuleContext(ListExprContext.class,0);
		}
		public MapExprContext mapExpr() {
			return getRuleContext(MapExprContext.class,0);
		}
		public SetExprContext setExpr() {
			return getRuleContext(SetExprContext.class,0);
		}
		public TupleExprContext tupleExpr() {
			return getRuleContext(TupleExprContext.class,0);
		}
		public PrimaryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary; }
	}

	public final PrimaryContext primary() throws RecognitionException {
		PrimaryContext _localctx = new PrimaryContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_primary);
		try {
			setState(948);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,118,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(935);
				identifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(936);
				match(SELF);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(937);
				literal();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(938);
				match(LPAREN);
				setState(939);
				expr(0);
				setState(940);
				match(RPAREN);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(942);
				block();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(943);
				structExpr();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(944);
				listExpr();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(945);
				mapExpr();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(946);
				setExpr();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(947);
				tupleExpr();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LambdaExprContext extends ParserRuleContext {
		public List<TerminalNode> BIT_OR() { return getTokens(UnifiedParser.BIT_OR); }
		public TerminalNode BIT_OR(int i) {
			return getToken(UnifiedParser.BIT_OR, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode MOVE() { return getToken(UnifiedParser.MOVE, 0); }
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(UnifiedParser.ARROW, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public LambdaExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lambdaExpr; }
	}

	public final LambdaExprContext lambdaExpr() throws RecognitionException {
		LambdaExprContext _localctx = new LambdaExprContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_lambdaExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(951);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==MOVE) {
				{
				setState(950);
				match(MOVE);
				}
			}

			setState(953);
			match(BIT_OR);
			setState(955);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SELF || _la==Identifier) {
				{
				setState(954);
				paramList();
				}
			}

			setState(957);
			match(BIT_OR);
			setState(964);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,122,_ctx) ) {
			case 1:
				{
				setState(960);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ARROW) {
					{
					setState(958);
					match(ARROW);
					setState(959);
					type(0);
					}
				}

				setState(962);
				block();
				}
				break;
			case 2:
				{
				setState(963);
				expr(0);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AsyncExprContext extends ParserRuleContext {
		public TerminalNode ASYNC() { return getToken(UnifiedParser.ASYNC, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public AsyncExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asyncExpr; }
	}

	public final AsyncExprContext asyncExpr() throws RecognitionException {
		AsyncExprContext _localctx = new AsyncExprContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_asyncExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(966);
			match(ASYNC);
			setState(967);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructExprContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public FieldInitListContext fieldInitList() {
			return getRuleContext(FieldInitListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public StructExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structExpr; }
	}

	public final StructExprContext structExpr() throws RecognitionException {
		StructExprContext _localctx = new StructExprContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_structExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(969);
			identifier();
			setState(970);
			match(LBRACE);
			setState(971);
			fieldInitList();
			setState(972);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldInitListContext extends ParserRuleContext {
		public List<FieldInitContext> fieldInit() {
			return getRuleContexts(FieldInitContext.class);
		}
		public FieldInitContext fieldInit(int i) {
			return getRuleContext(FieldInitContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public FieldInitListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldInitList; }
	}

	public final FieldInitListContext fieldInitList() throws RecognitionException {
		FieldInitListContext _localctx = new FieldInitListContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_fieldInitList);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(974);
			fieldInit();
			setState(979);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,123,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(975);
					match(COMMA);
					setState(976);
					fieldInit();
					}
					} 
				}
				setState(981);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,123,_ctx);
			}
			setState(983);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(982);
				match(COMMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldInitContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode RANGE() { return getToken(UnifiedParser.RANGE, 0); }
		public FieldInitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldInit; }
	}

	public final FieldInitContext fieldInit() throws RecognitionException {
		FieldInitContext _localctx = new FieldInitContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_fieldInit);
		try {
			setState(992);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,125,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(985);
				identifier();
				setState(986);
				match(COLON);
				setState(987);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(989);
				identifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(990);
				match(RANGE);
				setState(991);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListExprContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(UnifiedParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(UnifiedParser.RBRACK, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public TerminalNode FOR() { return getToken(UnifiedParser.FOR, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode IN() { return getToken(UnifiedParser.IN, 0); }
		public TerminalNode IF() { return getToken(UnifiedParser.IF, 0); }
		public ListExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listExpr; }
	}

	public final ListExprContext listExpr() throws RecognitionException {
		ListExprContext _localctx = new ListExprContext(_ctx, getState());
		enterRule(_localctx, 120, RULE_listExpr);
		int _la;
		try {
			int _alt;
			setState(1021);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,130,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(994);
				match(LBRACK);
				setState(1003);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
					{
					setState(995);
					expr(0);
					setState(1000);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,126,_ctx);
					while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
						if ( _alt==1 ) {
							{
							{
							setState(996);
							match(COMMA);
							setState(997);
							expr(0);
							}
							} 
						}
						setState(1002);
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,126,_ctx);
					}
					}
				}

				setState(1006);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(1005);
					match(COMMA);
					}
				}

				setState(1008);
				match(RBRACK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1009);
				match(LBRACK);
				setState(1010);
				expr(0);
				setState(1011);
				match(FOR);
				setState(1012);
				identifier();
				setState(1013);
				match(IN);
				setState(1014);
				expr(0);
				setState(1017);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IF) {
					{
					setState(1015);
					match(IF);
					setState(1016);
					expr(0);
					}
				}

				setState(1019);
				match(RBRACK);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MapExprContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<KeyValueContext> keyValue() {
			return getRuleContexts(KeyValueContext.class);
		}
		public KeyValueContext keyValue(int i) {
			return getRuleContext(KeyValueContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public MapExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapExpr; }
	}

	public final MapExprContext mapExpr() throws RecognitionException {
		MapExprContext _localctx = new MapExprContext(_ctx, getState());
		enterRule(_localctx, 122, RULE_mapExpr);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1023);
			match(LBRACE);
			setState(1032);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
				{
				setState(1024);
				keyValue();
				setState(1029);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,131,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1025);
						match(COMMA);
						setState(1026);
						keyValue();
						}
						} 
					}
					setState(1031);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,131,_ctx);
				}
				}
			}

			setState(1035);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(1034);
				match(COMMA);
				}
			}

			setState(1037);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeyValueContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public KeyValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyValue; }
	}

	public final KeyValueContext keyValue() throws RecognitionException {
		KeyValueContext _localctx = new KeyValueContext(_ctx, getState());
		enterRule(_localctx, 124, RULE_keyValue);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1039);
			expr(0);
			setState(1040);
			match(COLON);
			setState(1041);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SetExprContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public SetExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setExpr; }
	}

	public final SetExprContext setExpr() throws RecognitionException {
		SetExprContext _localctx = new SetExprContext(_ctx, getState());
		enterRule(_localctx, 126, RULE_setExpr);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1043);
			match(LBRACE);
			setState(1052);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
				{
				setState(1044);
				expr(0);
				setState(1049);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,134,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1045);
						match(COMMA);
						setState(1046);
						expr(0);
						}
						} 
					}
					setState(1051);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,134,_ctx);
				}
				}
			}

			setState(1055);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(1054);
				match(COMMA);
				}
			}

			setState(1057);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TupleExprContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(UnifiedParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(UnifiedParser.RPAREN, 0); }
		public List<NamedTupleFieldContext> namedTupleField() {
			return getRuleContexts(NamedTupleFieldContext.class);
		}
		public NamedTupleFieldContext namedTupleField(int i) {
			return getRuleContext(NamedTupleFieldContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public TupleExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleExpr; }
	}

	public final TupleExprContext tupleExpr() throws RecognitionException {
		TupleExprContext _localctx = new TupleExprContext(_ctx, getState());
		enterRule(_localctx, 128, RULE_tupleExpr);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1059);
			match(LPAREN);
			setState(1068);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
				{
				setState(1060);
				namedTupleField();
				setState(1065);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,137,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1061);
						match(COMMA);
						setState(1062);
						namedTupleField();
						}
						} 
					}
					setState(1067);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,137,_ctx);
				}
				}
			}

			setState(1071);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(1070);
				match(COMMA);
				}
			}

			setState(1073);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NamedTupleFieldContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode COLON() { return getToken(UnifiedParser.COLON, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NamedTupleFieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namedTupleField; }
	}

	public final NamedTupleFieldContext namedTupleField() throws RecognitionException {
		NamedTupleFieldContext _localctx = new NamedTupleFieldContext(_ctx, getState());
		enterRule(_localctx, 130, RULE_namedTupleField);
		try {
			setState(1080);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,140,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1075);
				identifier();
				setState(1076);
				match(COLON);
				setState(1077);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1079);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(UnifiedParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(UnifiedParser.RBRACE, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 132, RULE_block);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1082);
			match(LBRACE);
			setState(1086);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,141,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1083);
					statement();
					}
					} 
				}
				setState(1088);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,141,_ctx);
			}
			setState(1090);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688849930018858L) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 33350659L) != 0)) {
				{
				setState(1089);
				expr(0);
				}
			}

			setState(1092);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArgListContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(UnifiedParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(UnifiedParser.COMMA, i);
		}
		public ArgListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argList; }
	}

	public final ArgListContext argList() throws RecognitionException {
		ArgListContext _localctx = new ArgListContext(_ctx, getState());
		enterRule(_localctx, 134, RULE_argList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1094);
			expr(0);
			setState(1099);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(1095);
				match(COMMA);
				setState(1096);
				expr(0);
				}
				}
				setState(1101);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode IntLiteral() { return getToken(UnifiedParser.IntLiteral, 0); }
		public TerminalNode FloatLiteral() { return getToken(UnifiedParser.FloatLiteral, 0); }
		public TerminalNode StringLiteral() { return getToken(UnifiedParser.StringLiteral, 0); }
		public TerminalNode CharLiteral() { return getToken(UnifiedParser.CharLiteral, 0); }
		public TerminalNode BoolLiteral() { return getToken(UnifiedParser.BoolLiteral, 0); }
		public TerminalNode NullLiteral() { return getToken(UnifiedParser.NullLiteral, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 136, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1102);
			_la = _input.LA(1);
			if ( !(((((_la - 86)) & ~0x3f) == 0 && ((1L << (_la - 86)) & 123L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(UnifiedParser.Identifier, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 138, RULE_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1104);
			match(Identifier);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 29:
			return type_sempred((TypeContext)_localctx, predIndex);
		case 48:
			return pattern_sempred((PatternContext)_localctx, predIndex);
		case 52:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean type_sempred(TypeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pattern_sempred(PatternContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 18);
		case 4:
			return precpred(_ctx, 17);
		case 5:
			return precpred(_ctx, 16);
		case 6:
			return precpred(_ctx, 15);
		case 7:
			return precpred(_ctx, 14);
		case 8:
			return precpred(_ctx, 13);
		case 9:
			return precpred(_ctx, 12);
		case 10:
			return precpred(_ctx, 11);
		case 11:
			return precpred(_ctx, 10);
		case 12:
			return precpred(_ctx, 9);
		case 13:
			return precpred(_ctx, 7);
		case 14:
			return precpred(_ctx, 6);
		case 15:
			return precpred(_ctx, 5);
		case 16:
			return precpred(_ctx, 26);
		case 17:
			return precpred(_ctx, 25);
		case 18:
			return precpred(_ctx, 24);
		case 19:
			return precpred(_ctx, 23);
		case 20:
			return precpred(_ctx, 22);
		case 21:
			return precpred(_ctx, 8);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001`\u0453\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u00070\u00021\u00071\u0002"+
		"2\u00072\u00023\u00073\u00024\u00074\u00025\u00075\u00026\u00076\u0002"+
		"7\u00077\u00028\u00078\u00029\u00079\u0002:\u0007:\u0002;\u0007;\u0002"+
		"<\u0007<\u0002=\u0007=\u0002>\u0007>\u0002?\u0007?\u0002@\u0007@\u0002"+
		"A\u0007A\u0002B\u0007B\u0002C\u0007C\u0002D\u0007D\u0002E\u0007E\u0001"+
		"\u0000\u0005\u0000\u008e\b\u0000\n\u0000\f\u0000\u0091\t\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\u009d\b\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002\u00a3\b\u0002\n\u0002"+
		"\f\u0002\u00a6\t\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0003\u0003\u00ae\b\u0003\u0001\u0003\u0003\u0003"+
		"\u00b1\b\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0003\u0003\u00b9\b\u0003\u0003\u0003\u00bb\b\u0003\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0005\u0004\u00c0\b\u0004\n\u0004\f\u0004"+
		"\u00c3\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00c8\b"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00ce"+
		"\b\u0005\u0005\u0005\u00d0\b\u0005\n\u0005\f\u0005\u00d3\t\u0005\u0001"+
		"\u0005\u0003\u0005\u00d6\b\u0005\u0001\u0006\u0003\u0006\u00d9\b\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u00de\b\u0006\u0001\u0006"+
		"\u0001\u0006\u0003\u0006\u00e2\b\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0003\u0006\u00e7\b\u0006\u0001\u0006\u0003\u0006\u00ea\b\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u00f1"+
		"\b\u0007\n\u0007\f\u0007\u00f4\t\u0007\u0001\b\u0001\b\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\b\u0104\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0005\t\u010a"+
		"\b\t\n\t\f\t\u010d\t\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0003\n"+
		"\u0114\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u0119\b\u000b"+
		"\n\u000b\f\u000b\u011c\t\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f"+
		"\u0122\b\f\n\f\f\f\u0125\t\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0003\u000e\u012c\b\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e"+
		"\u0131\b\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u0135\b\u000e\n\u000e"+
		"\f\u000e\u0138\t\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0003\u000f"+
		"\u013d\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0003\u000f\u0144\b\u000f\u0001\u0010\u0003\u0010\u0147\b\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u014c\b\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u0152\b\u0010\n\u0010\f\u0010"+
		"\u0155\t\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0003\u0011\u015e\b\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0005\u0012\u0163\b\u0012\n\u0012\f\u0012\u0166\t\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u016b\b\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0014\u0003\u0014\u0170\b\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0003\u0014\u0175\b\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u0179"+
		"\b\u0014\n\u0014\f\u0014\u017c\t\u0014\u0001\u0014\u0001\u0014\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0184\b\u0015\u0003\u0015"+
		"\u0186\b\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u018b\b"+
		"\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u018f\b\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0003\u0016\u0194\b\u0016\u0001\u0016\u0003\u0016\u0197"+
		"\b\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u019c\b\u0016"+
		"\u0001\u0016\u0001\u0016\u0003\u0016\u01a0\b\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0003\u0016\u01a5\b\u0016\u0001\u0016\u0001\u0016\u0003\u0016"+
		"\u01a9\b\u0016\u0001\u0017\u0001\u0017\u0003\u0017\u01ad\b\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u01b3\b\u0017\u0001"+
		"\u0017\u0001\u0017\u0005\u0017\u01b7\b\u0017\n\u0017\f\u0017\u01ba\t\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u01c0\b\u0017"+
		"\u0001\u0017\u0001\u0017\u0003\u0017\u01c4\b\u0017\u0001\u0017\u0001\u0017"+
		"\u0005\u0017\u01c8\b\u0017\n\u0017\f\u0017\u01cb\t\u0017\u0001\u0017\u0001"+
		"\u0017\u0003\u0017\u01cf\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u01d7\b\u0018\u0003\u0018\u01d9"+
		"\b\u0018\u0001\u0019\u0003\u0019\u01dc\b\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0003\u0019\u01e1\b\u0019\u0001\u0019\u0001\u0019\u0005\u0019"+
		"\u01e5\b\u0019\n\u0019\f\u0019\u01e8\t\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u001a\u0003\u001a\u01ed\b\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0003\u001a\u01f5\b\u001a\u0001\u001a\u0003"+
		"\u001a\u01f8\b\u001a\u0001\u001a\u0003\u001a\u01fb\b\u001a\u0001\u001b"+
		"\u0003\u001b\u01fe\b\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b"+
		"\u0203\b\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u0208\b"+
		"\u001b\u0001\u001b\u0003\u001b\u020b\b\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0003\u001b\u0210\b\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u021a"+
		"\b\u001b\u0003\u001b\u021c\b\u001b\u0001\u001c\u0003\u001c\u021f\b\u001c"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0003\u001c\u0228\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0003\u001d\u0230\b\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0003\u001d\u0235\b\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0003\u001d\u0245\b\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0004\u001d\u024d\b\u001d\u000b\u001d\f\u001d"+
		"\u024e\u0005\u001d\u0251\b\u001d\n\u001d\f\u001d\u0254\t\u001d\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0005\u001e\u0259\b\u001e\n\u001e\f\u001e\u025c"+
		"\t\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u026c\b\u001f\u0001 \u0001"+
		" \u0003 \u0270\b \u0001 \u0001 \u0001 \u0003 \u0275\b \u0001 \u0001 \u0001"+
		" \u0003 \u027a\b \u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0003!\u0282"+
		"\b!\u0001!\u0003!\u0285\b!\u0001\"\u0001\"\u0001\"\u0001\"\u0003\"\u028b"+
		"\b\"\u0001#\u0001#\u0001$\u0001$\u0001$\u0001$\u0001%\u0001%\u0003%\u0295"+
		"\b%\u0001&\u0001&\u0003&\u0299\b&\u0001&\u0003&\u029c\b&\u0001\'\u0001"+
		"\'\u0001\'\u0001\'\u0001\'\u0001\'\u0001\'\u0001\'\u0005\'\u02a6\b\'\n"+
		"\'\f\'\u02a9\t\'\u0001\'\u0001\'\u0003\'\u02ad\b\'\u0001(\u0001(\u0001"+
		"(\u0003(\u02b2\b(\u0001(\u0001(\u0001(\u0001)\u0001)\u0001)\u0003)\u02ba"+
		"\b)\u0001)\u0001)\u0001)\u0001)\u0001*\u0001*\u0001*\u0003*\u02c3\b*\u0001"+
		"*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001+\u0001+\u0001+\u0001+\u0005"+
		"+\u02cf\b+\n+\f+\u02d2\t+\u0001+\u0001+\u0001,\u0001,\u0001,\u0001,\u0001"+
		",\u0003,\u02db\b,\u0001-\u0001-\u0003-\u02df\b-\u0001-\u0003-\u02e2\b"+
		"-\u0001.\u0001.\u0003.\u02e6\b.\u0001.\u0003.\u02e9\b.\u0001/\u0001/\u0001"+
		"0\u00010\u00010\u00010\u00010\u00010\u00010\u00030\u02f4\b0\u00010\u0001"+
		"0\u00010\u00010\u00010\u00010\u00030\u02fc\b0\u00010\u00010\u00030\u0300"+
		"\b0\u00010\u00010\u00010\u00010\u00010\u00010\u00010\u00030\u0309\b0\u0001"+
		"0\u00030\u030c\b0\u00010\u00010\u00010\u00050\u0311\b0\n0\f0\u0314\t0"+
		"\u00011\u00011\u00011\u00051\u0319\b1\n1\f1\u031c\t1\u00012\u00012\u0001"+
		"2\u00012\u00012\u00012\u00032\u0324\b2\u00013\u00013\u00013\u00053\u0329"+
		"\b3\n3\f3\u032c\t3\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u0001"+
		"4\u00054\u0340\b4\n4\f4\u0343\t4\u00014\u00014\u00034\u0347\b4\u00014"+
		"\u00014\u00014\u00014\u00054\u034d\b4\n4\f4\u0350\t4\u00014\u00014\u0003"+
		"4\u0354\b4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00034\u0384\b4\u0001"+
		"4\u00014\u00014\u00014\u00014\u00034\u038b\b4\u00014\u00014\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00034\u0397\b4\u00014\u0001"+
		"4\u00014\u00014\u00014\u00054\u039e\b4\n4\f4\u03a1\t4\u00015\u00015\u0001"+
		"5\u00015\u00015\u00016\u00016\u00016\u00016\u00016\u00016\u00016\u0001"+
		"6\u00016\u00016\u00016\u00016\u00016\u00036\u03b5\b6\u00017\u00037\u03b8"+
		"\b7\u00017\u00017\u00037\u03bc\b7\u00017\u00017\u00017\u00037\u03c1\b"+
		"7\u00017\u00017\u00037\u03c5\b7\u00018\u00018\u00018\u00019\u00019\u0001"+
		"9\u00019\u00019\u0001:\u0001:\u0001:\u0005:\u03d2\b:\n:\f:\u03d5\t:\u0001"+
		":\u0003:\u03d8\b:\u0001;\u0001;\u0001;\u0001;\u0001;\u0001;\u0001;\u0003"+
		";\u03e1\b;\u0001<\u0001<\u0001<\u0001<\u0005<\u03e7\b<\n<\f<\u03ea\t<"+
		"\u0003<\u03ec\b<\u0001<\u0003<\u03ef\b<\u0001<\u0001<\u0001<\u0001<\u0001"+
		"<\u0001<\u0001<\u0001<\u0001<\u0003<\u03fa\b<\u0001<\u0001<\u0003<\u03fe"+
		"\b<\u0001=\u0001=\u0001=\u0001=\u0005=\u0404\b=\n=\f=\u0407\t=\u0003="+
		"\u0409\b=\u0001=\u0003=\u040c\b=\u0001=\u0001=\u0001>\u0001>\u0001>\u0001"+
		">\u0001?\u0001?\u0001?\u0001?\u0005?\u0418\b?\n?\f?\u041b\t?\u0003?\u041d"+
		"\b?\u0001?\u0003?\u0420\b?\u0001?\u0001?\u0001@\u0001@\u0001@\u0001@\u0005"+
		"@\u0428\b@\n@\f@\u042b\t@\u0003@\u042d\b@\u0001@\u0003@\u0430\b@\u0001"+
		"@\u0001@\u0001A\u0001A\u0001A\u0001A\u0001A\u0003A\u0439\bA\u0001B\u0001"+
		"B\u0005B\u043d\bB\nB\fB\u0440\tB\u0001B\u0003B\u0443\bB\u0001B\u0001B"+
		"\u0001C\u0001C\u0001C\u0005C\u044a\bC\nC\fC\u044d\tC\u0001D\u0001D\u0001"+
		"E\u0001E\u0001E\u0000\u0003:`hF\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPR"+
		"TVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084\u0086\u0088\u008a\u0000\u000b"+
		"\u0001\u0000!&\u0002\u0000\b\b\u000b\u000b\u0004\u0000\r\u000e\u0015\u0015"+
		"\u001a\u001a12\u0001\u0000\u000f\u0011\u0001\u0000\r\u000e\u0001\u0000"+
		"\u0016\u0017\u0001\u0000\u001d \u0001\u0000\u001b\u001c\u0002\u0000!+"+
		".0\u0001\u000012\u0002\u0000VWY\\\u04dc\u0000\u008f\u0001\u0000\u0000"+
		"\u0000\u0002\u009c\u0001\u0000\u0000\u0000\u0004\u009e\u0001\u0000\u0000"+
		"\u0000\u0006\u00ba\u0001\u0000\u0000\u0000\b\u00bc\u0001\u0000\u0000\u0000"+
		"\n\u00d5\u0001\u0000\u0000\u0000\f\u00d8\u0001\u0000\u0000\u0000\u000e"+
		"\u00ed\u0001\u0000\u0000\u0000\u0010\u0103\u0001\u0000\u0000\u0000\u0012"+
		"\u0105\u0001\u0000\u0000\u0000\u0014\u0110\u0001\u0000\u0000\u0000\u0016"+
		"\u0115\u0001\u0000\u0000\u0000\u0018\u011d\u0001\u0000\u0000\u0000\u001a"+
		"\u0126\u0001\u0000\u0000\u0000\u001c\u012b\u0001\u0000\u0000\u0000\u001e"+
		"\u0143\u0001\u0000\u0000\u0000 \u0146\u0001\u0000\u0000\u0000\"\u0158"+
		"\u0001\u0000\u0000\u0000$\u015f\u0001\u0000\u0000\u0000&\u016a\u0001\u0000"+
		"\u0000\u0000(\u016f\u0001\u0000\u0000\u0000*\u0185\u0001\u0000\u0000\u0000"+
		",\u01a8\u0001\u0000\u0000\u0000.\u01ce\u0001\u0000\u0000\u00000\u01d8"+
		"\u0001\u0000\u0000\u00002\u01db\u0001\u0000\u0000\u00004\u01fa\u0001\u0000"+
		"\u0000\u00006\u021b\u0001\u0000\u0000\u00008\u021e\u0001\u0000\u0000\u0000"+
		":\u0244\u0001\u0000\u0000\u0000<\u0255\u0001\u0000\u0000\u0000>\u026b"+
		"\u0001\u0000\u0000\u0000@\u026d\u0001\u0000\u0000\u0000B\u027b\u0001\u0000"+
		"\u0000\u0000D\u0286\u0001\u0000\u0000\u0000F\u028c\u0001\u0000\u0000\u0000"+
		"H\u028e\u0001\u0000\u0000\u0000J\u0292\u0001\u0000\u0000\u0000L\u0296"+
		"\u0001\u0000\u0000\u0000N\u029d\u0001\u0000\u0000\u0000P\u02b1\u0001\u0000"+
		"\u0000\u0000R\u02b9\u0001\u0000\u0000\u0000T\u02c2\u0001\u0000\u0000\u0000"+
		"V\u02ca\u0001\u0000\u0000\u0000X\u02d5\u0001\u0000\u0000\u0000Z\u02dc"+
		"\u0001\u0000\u0000\u0000\\\u02e3\u0001\u0000\u0000\u0000^\u02ea\u0001"+
		"\u0000\u0000\u0000`\u030b\u0001\u0000\u0000\u0000b\u0315\u0001\u0000\u0000"+
		"\u0000d\u0323\u0001\u0000\u0000\u0000f\u0325\u0001\u0000\u0000\u0000h"+
		"\u0353\u0001\u0000\u0000\u0000j\u03a2\u0001\u0000\u0000\u0000l\u03b4\u0001"+
		"\u0000\u0000\u0000n\u03b7\u0001\u0000\u0000\u0000p\u03c6\u0001\u0000\u0000"+
		"\u0000r\u03c9\u0001\u0000\u0000\u0000t\u03ce\u0001\u0000\u0000\u0000v"+
		"\u03e0\u0001\u0000\u0000\u0000x\u03fd\u0001\u0000\u0000\u0000z\u03ff\u0001"+
		"\u0000\u0000\u0000|\u040f\u0001\u0000\u0000\u0000~\u0413\u0001\u0000\u0000"+
		"\u0000\u0080\u0423\u0001\u0000\u0000\u0000\u0082\u0438\u0001\u0000\u0000"+
		"\u0000\u0084\u043a\u0001\u0000\u0000\u0000\u0086\u0446\u0001\u0000\u0000"+
		"\u0000\u0088\u044e\u0001\u0000\u0000\u0000\u008a\u0450\u0001\u0000\u0000"+
		"\u0000\u008c\u008e\u0003\u0002\u0001\u0000\u008d\u008c\u0001\u0000\u0000"+
		"\u0000\u008e\u0091\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000"+
		"\u0000\u008f\u0090\u0001\u0000\u0000\u0000\u0090\u0001\u0001\u0000\u0000"+
		"\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0092\u009d\u0003\u0004\u0002"+
		"\u0000\u0093\u009d\u0003\f\u0006\u0000\u0094\u009d\u0003\u001c\u000e\u0000"+
		"\u0095\u009d\u0003 \u0010\u0000\u0096\u009d\u0003(\u0014\u0000\u0097\u009d"+
		"\u0003.\u0017\u0000\u0098\u009d\u00032\u0019\u0000\u0099\u009d\u00036"+
		"\u001b\u0000\u009a\u009d\u0003\u0006\u0003\u0000\u009b\u009d\u00038\u001c"+
		"\u0000\u009c\u0092\u0001\u0000\u0000\u0000\u009c\u0093\u0001\u0000\u0000"+
		"\u0000\u009c\u0094\u0001\u0000\u0000\u0000\u009c\u0095\u0001\u0000\u0000"+
		"\u0000\u009c\u0096\u0001\u0000\u0000\u0000\u009c\u0097\u0001\u0000\u0000"+
		"\u0000\u009c\u0098\u0001\u0000\u0000\u0000\u009c\u0099\u0001\u0000\u0000"+
		"\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009c\u009b\u0001\u0000\u0000"+
		"\u0000\u009d\u0003\u0001\u0000\u0000\u0000\u009e\u009f\u00054\u0000\u0000"+
		"\u009f\u00a0\u0003\u008aE\u0000\u00a0\u00a4\u0005\u0001\u0000\u0000\u00a1"+
		"\u00a3\u0003\u0002\u0001\u0000\u00a2\u00a1\u0001\u0000\u0000\u0000\u00a3"+
		"\u00a6\u0001\u0000\u0000\u0000\u00a4\u00a2\u0001\u0000\u0000\u0000\u00a4"+
		"\u00a5\u0001\u0000\u0000\u0000\u00a5\u00a7\u0001\u0000\u0000\u0000\u00a6"+
		"\u00a4\u0001\u0000\u0000\u0000\u00a7\u00a8\u0005\u0002\u0000\u0000\u00a8"+
		"\u0005\u0001\u0000\u0000\u0000\u00a9\u00aa\u00055\u0000\u0000\u00aa\u00ad"+
		"\u0003\b\u0004\u0000\u00ab\u00ac\u00056\u0000\u0000\u00ac\u00ae\u0003"+
		"\u008aE\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000"+
		"\u0000\u0000\u00ae\u00b0\u0001\u0000\u0000\u0000\u00af\u00b1\u0005\u0007"+
		"\u0000\u0000\u00b0\u00af\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001\u0000"+
		"\u0000\u0000\u00b1\u00bb\u0001\u0000\u0000\u0000\u00b2\u00b3\u00055\u0000"+
		"\u0000\u00b3\u00b4\u0003\b\u0004\u0000\u00b4\u00b5\u0005\u0001\u0000\u0000"+
		"\u00b5\u00b6\u0003\n\u0005\u0000\u00b6\u00b8\u0005\u0002\u0000\u0000\u00b7"+
		"\u00b9\u0005\u0007\u0000\u0000\u00b8\u00b7\u0001\u0000\u0000\u0000\u00b8"+
		"\u00b9\u0001\u0000\u0000\u0000\u00b9\u00bb\u0001\u0000\u0000\u0000\u00ba"+
		"\u00a9\u0001\u0000\u0000\u0000\u00ba\u00b2\u0001\u0000\u0000\u0000\u00bb"+
		"\u0007\u0001\u0000\u0000\u0000\u00bc\u00c1\u0003\u008aE\u0000\u00bd\u00be"+
		"\u0005\n\u0000\u0000\u00be\u00c0\u0003\u008aE\u0000\u00bf\u00bd\u0001"+
		"\u0000\u0000\u0000\u00c0\u00c3\u0001\u0000\u0000\u0000\u00c1\u00bf\u0001"+
		"\u0000\u0000\u0000\u00c1\u00c2\u0001\u0000\u0000\u0000\u00c2\t\u0001\u0000"+
		"\u0000\u0000\u00c3\u00c1\u0001\u0000\u0000\u0000\u00c4\u00c7\u0003\u008a"+
		"E\u0000\u00c5\u00c6\u00056\u0000\u0000\u00c6\u00c8\u0003\u008aE\u0000"+
		"\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c7\u00c8\u0001\u0000\u0000\u0000"+
		"\u00c8\u00d1\u0001\u0000\u0000\u0000\u00c9\u00ca\u0005\t\u0000\u0000\u00ca"+
		"\u00cd\u0003\u008aE\u0000\u00cb\u00cc\u00056\u0000\u0000\u00cc\u00ce\u0003"+
		"\u008aE\u0000\u00cd\u00cb\u0001\u0000\u0000\u0000\u00cd\u00ce\u0001\u0000"+
		"\u0000\u0000\u00ce\u00d0\u0001\u0000\u0000\u0000\u00cf\u00c9\u0001\u0000"+
		"\u0000\u0000\u00d0\u00d3\u0001\u0000\u0000\u0000\u00d1\u00cf\u0001\u0000"+
		"\u0000\u0000\u00d1\u00d2\u0001\u0000\u0000\u0000\u00d2\u00d6\u0001\u0000"+
		"\u0000\u0000\u00d3\u00d1\u0001\u0000\u0000\u0000\u00d4\u00d6\u0005\u000f"+
		"\u0000\u0000\u00d5\u00c4\u0001\u0000\u0000\u0000\u00d5\u00d4\u0001\u0000"+
		"\u0000\u0000\u00d6\u000b\u0001\u0000\u0000\u0000\u00d7\u00d9\u0005C\u0000"+
		"\u0000\u00d8\u00d7\u0001\u0000\u0000\u0000\u00d8\u00d9\u0001\u0000\u0000"+
		"\u0000\u00d9\u00da\u0001\u0000\u0000\u0000\u00da\u00db\u00057\u0000\u0000"+
		"\u00db\u00dd\u0003\u008aE\u0000\u00dc\u00de\u0003\u0012\t\u0000\u00dd"+
		"\u00dc\u0001\u0000\u0000\u0000\u00dd\u00de\u0001\u0000\u0000\u0000\u00de"+
		"\u00df\u0001\u0000\u0000\u0000\u00df\u00e1\u0005\u0003\u0000\u0000\u00e0"+
		"\u00e2\u0003\u000e\u0007\u0000\u00e1\u00e0\u0001\u0000\u0000\u0000\u00e1"+
		"\u00e2\u0001\u0000\u0000\u0000\u00e2\u00e3\u0001\u0000\u0000\u0000\u00e3"+
		"\u00e6\u0005\u0004\u0000\u0000\u00e4\u00e5\u0005\u000b\u0000\u0000\u00e5"+
		"\u00e7\u0003:\u001d\u0000\u00e6\u00e4\u0001\u0000\u0000\u0000\u00e6\u00e7"+
		"\u0001\u0000\u0000\u0000\u00e7\u00e9\u0001\u0000\u0000\u0000\u00e8\u00ea"+
		"\u0003\u0018\f\u0000\u00e9\u00e8\u0001\u0000\u0000\u0000\u00e9\u00ea\u0001"+
		"\u0000\u0000\u0000\u00ea\u00eb\u0001\u0000\u0000\u0000\u00eb\u00ec\u0003"+
		"\u0084B\u0000\u00ec\r\u0001\u0000\u0000\u0000\u00ed\u00f2\u0003\u0010"+
		"\b\u0000\u00ee\u00ef\u0005\t\u0000\u0000\u00ef\u00f1\u0003\u0010\b\u0000"+
		"\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f1\u00f4\u0001\u0000\u0000\u0000"+
		"\u00f2\u00f0\u0001\u0000\u0000\u0000\u00f2\u00f3\u0001\u0000\u0000\u0000"+
		"\u00f3\u000f\u0001\u0000\u0000\u0000\u00f4\u00f2\u0001\u0000\u0000\u0000"+
		"\u00f5\u00f6\u0003\u008aE\u0000\u00f6\u00f7\u0005\b\u0000\u0000\u00f7"+
		"\u00f8\u0003:\u001d\u0000\u00f8\u0104\u0001\u0000\u0000\u0000\u00f9\u0104"+
		"\u0005D\u0000\u0000\u00fa\u00fb\u0005D\u0000\u0000\u00fb\u00fc\u0005\b"+
		"\u0000\u0000\u00fc\u00fd\u0005\u0012\u0000\u0000\u00fd\u0104\u0005D\u0000"+
		"\u0000\u00fe\u00ff\u0005D\u0000\u0000\u00ff\u0100\u0005\b\u0000\u0000"+
		"\u0100\u0101\u0005\u0012\u0000\u0000\u0101\u0102\u0005B\u0000\u0000\u0102"+
		"\u0104\u0005D\u0000\u0000\u0103\u00f5\u0001\u0000\u0000\u0000\u0103\u00f9"+
		"\u0001\u0000\u0000\u0000\u0103\u00fa\u0001\u0000\u0000\u0000\u0103\u00fe"+
		"\u0001\u0000\u0000\u0000\u0104\u0011\u0001\u0000\u0000\u0000\u0105\u0106"+
		"\u0005\u001d\u0000\u0000\u0106\u010b\u0003\u0014\n\u0000\u0107\u0108\u0005"+
		"\t\u0000\u0000\u0108\u010a\u0003\u0014\n\u0000\u0109\u0107\u0001\u0000"+
		"\u0000\u0000\u010a\u010d\u0001\u0000\u0000\u0000\u010b\u0109\u0001\u0000"+
		"\u0000\u0000\u010b\u010c\u0001\u0000\u0000\u0000\u010c\u010e\u0001\u0000"+
		"\u0000\u0000\u010d\u010b\u0001\u0000\u0000\u0000\u010e\u010f\u0005\u001e"+
		"\u0000\u0000\u010f\u0013\u0001\u0000\u0000\u0000\u0110\u0113\u0003\u008a"+
		"E\u0000\u0111\u0112\u0005\b\u0000\u0000\u0112\u0114\u0003\u0016\u000b"+
		"\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0113\u0114\u0001\u0000\u0000"+
		"\u0000\u0114\u0015\u0001\u0000\u0000\u0000\u0115\u011a\u0003:\u001d\u0000"+
		"\u0116\u0117\u0005\r\u0000\u0000\u0117\u0119\u0003:\u001d\u0000\u0118"+
		"\u0116\u0001\u0000\u0000\u0000\u0119\u011c\u0001\u0000\u0000\u0000\u011a"+
		"\u0118\u0001\u0000\u0000\u0000\u011a\u011b\u0001\u0000\u0000\u0000\u011b"+
		"\u0017\u0001\u0000\u0000\u0000\u011c\u011a\u0001\u0000\u0000\u0000\u011d"+
		"\u011e\u0005P\u0000\u0000\u011e\u0123\u0003\u001a\r\u0000\u011f\u0120"+
		"\u0005\t\u0000\u0000\u0120\u0122\u0003\u001a\r\u0000\u0121\u011f\u0001"+
		"\u0000\u0000\u0000\u0122\u0125\u0001\u0000\u0000\u0000\u0123\u0121\u0001"+
		"\u0000\u0000\u0000\u0123\u0124\u0001\u0000\u0000\u0000\u0124\u0019\u0001"+
		"\u0000\u0000\u0000\u0125\u0123\u0001\u0000\u0000\u0000\u0126\u0127\u0003"+
		":\u001d\u0000\u0127\u0128\u0005\b\u0000\u0000\u0128\u0129\u0003\u0016"+
		"\u000b\u0000\u0129\u001b\u0001\u0000\u0000\u0000\u012a\u012c\u0005C\u0000"+
		"\u0000\u012b\u012a\u0001\u0000\u0000\u0000\u012b\u012c\u0001\u0000\u0000"+
		"\u0000\u012c\u012d\u0001\u0000\u0000\u0000\u012d\u012e\u00058\u0000\u0000"+
		"\u012e\u0130\u0003\u008aE\u0000\u012f\u0131\u0003\u0012\t\u0000\u0130"+
		"\u012f\u0001\u0000\u0000\u0000\u0130\u0131\u0001\u0000\u0000\u0000\u0131"+
		"\u0132\u0001\u0000\u0000\u0000\u0132\u0136\u0005\u0001\u0000\u0000\u0133"+
		"\u0135\u0003\u001e\u000f\u0000\u0134\u0133\u0001\u0000\u0000\u0000\u0135"+
		"\u0138\u0001\u0000\u0000\u0000\u0136\u0134\u0001\u0000\u0000\u0000\u0136"+
		"\u0137\u0001\u0000\u0000\u0000\u0137\u0139\u0001\u0000\u0000\u0000\u0138"+
		"\u0136\u0001\u0000\u0000\u0000\u0139\u013a\u0005\u0002\u0000\u0000\u013a"+
		"\u001d\u0001\u0000\u0000\u0000\u013b\u013d\u0005C\u0000\u0000\u013c\u013b"+
		"\u0001\u0000\u0000\u0000\u013c\u013d\u0001\u0000\u0000\u0000\u013d\u013e"+
		"\u0001\u0000\u0000\u0000\u013e\u013f\u0003\u008aE\u0000\u013f\u0140\u0005"+
		"\b\u0000\u0000\u0140\u0141\u0003:\u001d\u0000\u0141\u0144\u0001\u0000"+
		"\u0000\u0000\u0142\u0144\u0003\f\u0006\u0000\u0143\u013c\u0001\u0000\u0000"+
		"\u0000\u0143\u0142\u0001\u0000\u0000\u0000\u0144\u001f\u0001\u0000\u0000"+
		"\u0000\u0145\u0147\u0005C\u0000\u0000\u0146\u0145\u0001\u0000\u0000\u0000"+
		"\u0146\u0147\u0001\u0000\u0000\u0000\u0147\u0148\u0001\u0000\u0000\u0000"+
		"\u0148\u0149\u00059\u0000\u0000\u0149\u014b\u0003\u008aE\u0000\u014a\u014c"+
		"\u0003\u0012\t\u0000\u014b\u014a\u0001\u0000\u0000\u0000\u014b\u014c\u0001"+
		"\u0000\u0000\u0000\u014c\u014d\u0001\u0000\u0000\u0000\u014d\u014e\u0005"+
		"\u0001\u0000\u0000\u014e\u0153\u0003\"\u0011\u0000\u014f\u0150\u0005\t"+
		"\u0000\u0000\u0150\u0152\u0003\"\u0011\u0000\u0151\u014f\u0001\u0000\u0000"+
		"\u0000\u0152\u0155\u0001\u0000\u0000\u0000\u0153\u0151\u0001\u0000\u0000"+
		"\u0000\u0153\u0154\u0001\u0000\u0000\u0000\u0154\u0156\u0001\u0000\u0000"+
		"\u0000\u0155\u0153\u0001\u0000\u0000\u0000\u0156\u0157\u0005\u0002\u0000"+
		"\u0000\u0157!\u0001\u0000\u0000\u0000\u0158\u015d\u0003\u008aE\u0000\u0159"+
		"\u015a\u0005\u0003\u0000\u0000\u015a\u015b\u0003$\u0012\u0000\u015b\u015c"+
		"\u0005\u0004\u0000\u0000\u015c\u015e\u0001\u0000\u0000\u0000\u015d\u0159"+
		"\u0001\u0000\u0000\u0000\u015d\u015e\u0001\u0000\u0000\u0000\u015e#\u0001"+
		"\u0000\u0000\u0000\u015f\u0164\u0003&\u0013\u0000\u0160\u0161\u0005\t"+
		"\u0000\u0000\u0161\u0163\u0003&\u0013\u0000\u0162\u0160\u0001\u0000\u0000"+
		"\u0000\u0163\u0166\u0001\u0000\u0000\u0000\u0164\u0162\u0001\u0000\u0000"+
		"\u0000\u0164\u0165\u0001\u0000\u0000\u0000\u0165%\u0001\u0000\u0000\u0000"+
		"\u0166\u0164\u0001\u0000\u0000\u0000\u0167\u0168\u0003\u008aE\u0000\u0168"+
		"\u0169\u0005\b\u0000\u0000\u0169\u016b\u0001\u0000\u0000\u0000\u016a\u0167"+
		"\u0001\u0000\u0000\u0000\u016a\u016b\u0001\u0000\u0000\u0000\u016b\u016c"+
		"\u0001\u0000\u0000\u0000\u016c\u016d\u0003:\u001d\u0000\u016d\'\u0001"+
		"\u0000\u0000\u0000\u016e\u0170\u0005C\u0000\u0000\u016f\u016e\u0001\u0000"+
		"\u0000\u0000\u016f\u0170\u0001\u0000\u0000\u0000\u0170\u0171\u0001\u0000"+
		"\u0000\u0000\u0171\u0172\u0005:\u0000\u0000\u0172\u0174\u0003\u008aE\u0000"+
		"\u0173\u0175\u0003\u0012\t\u0000\u0174\u0173\u0001\u0000\u0000\u0000\u0174"+
		"\u0175\u0001\u0000\u0000\u0000\u0175\u0176\u0001\u0000\u0000\u0000\u0176"+
		"\u017a\u0005\u0001\u0000\u0000\u0177\u0179\u0003*\u0015\u0000\u0178\u0177"+
		"\u0001\u0000\u0000\u0000\u0179\u017c\u0001\u0000\u0000\u0000\u017a\u0178"+
		"\u0001\u0000\u0000\u0000\u017a\u017b\u0001\u0000\u0000\u0000\u017b\u017d"+
		"\u0001\u0000\u0000\u0000\u017c\u017a\u0001\u0000\u0000\u0000\u017d\u017e"+
		"\u0005\u0002\u0000\u0000\u017e)\u0001\u0000\u0000\u0000\u017f\u0186\u0003"+
		",\u0016\u0000\u0180\u0181\u0005=\u0000\u0000\u0181\u0183\u0003\u008aE"+
		"\u0000\u0182\u0184\u0005\u0007\u0000\u0000\u0183\u0182\u0001\u0000\u0000"+
		"\u0000\u0183\u0184\u0001\u0000\u0000\u0000\u0184\u0186\u0001\u0000\u0000"+
		"\u0000\u0185\u017f\u0001\u0000\u0000\u0000\u0185\u0180\u0001\u0000\u0000"+
		"\u0000\u0186+\u0001\u0000\u0000\u0000\u0187\u0188\u00057\u0000\u0000\u0188"+
		"\u018a\u0003\u008aE\u0000\u0189\u018b\u0003\u0012\t\u0000\u018a\u0189"+
		"\u0001\u0000\u0000\u0000\u018a\u018b\u0001\u0000\u0000\u0000\u018b\u018c"+
		"\u0001\u0000\u0000\u0000\u018c\u018e\u0005\u0003\u0000\u0000\u018d\u018f"+
		"\u0003\u000e\u0007\u0000\u018e\u018d\u0001\u0000\u0000\u0000\u018e\u018f"+
		"\u0001\u0000\u0000\u0000\u018f\u0190\u0001\u0000\u0000\u0000\u0190\u0193"+
		"\u0005\u0004\u0000\u0000\u0191\u0192\u0005\u000b\u0000\u0000\u0192\u0194"+
		"\u0003:\u001d\u0000\u0193\u0191\u0001\u0000\u0000\u0000\u0193\u0194\u0001"+
		"\u0000\u0000\u0000\u0194\u0196\u0001\u0000\u0000\u0000\u0195\u0197\u0005"+
		"\u0007\u0000\u0000\u0196\u0195\u0001\u0000\u0000\u0000\u0196\u0197\u0001"+
		"\u0000\u0000\u0000\u0197\u01a9\u0001\u0000\u0000\u0000\u0198\u0199\u0005"+
		"7\u0000\u0000\u0199\u019b\u0003\u008aE\u0000\u019a\u019c\u0003\u0012\t"+
		"\u0000\u019b\u019a\u0001\u0000\u0000\u0000\u019b\u019c\u0001\u0000\u0000"+
		"\u0000\u019c\u019d\u0001\u0000\u0000\u0000\u019d\u019f\u0005\u0003\u0000"+
		"\u0000\u019e\u01a0\u0003\u000e\u0007\u0000\u019f\u019e\u0001\u0000\u0000"+
		"\u0000\u019f\u01a0\u0001\u0000\u0000\u0000\u01a0\u01a1\u0001\u0000\u0000"+
		"\u0000\u01a1\u01a4\u0005\u0004\u0000\u0000\u01a2\u01a3\u0005\u000b\u0000"+
		"\u0000\u01a3\u01a5\u0003:\u001d\u0000\u01a4\u01a2\u0001\u0000\u0000\u0000"+
		"\u01a4\u01a5\u0001\u0000\u0000\u0000\u01a5\u01a6\u0001\u0000\u0000\u0000"+
		"\u01a6\u01a7\u0003\u0084B\u0000\u01a7\u01a9\u0001\u0000\u0000\u0000\u01a8"+
		"\u0187\u0001\u0000\u0000\u0000\u01a8\u0198\u0001\u0000\u0000\u0000\u01a9"+
		"-\u0001\u0000\u0000\u0000\u01aa\u01ac\u0005;\u0000\u0000\u01ab\u01ad\u0003"+
		"\u0012\t\u0000\u01ac\u01ab\u0001\u0000\u0000\u0000\u01ac\u01ad\u0001\u0000"+
		"\u0000\u0000\u01ad\u01ae\u0001\u0000\u0000\u0000\u01ae\u01af\u0003:\u001d"+
		"\u0000\u01af\u01b0\u0005I\u0000\u0000\u01b0\u01b2\u0003:\u001d\u0000\u01b1"+
		"\u01b3\u0003\u0018\f\u0000\u01b2\u01b1\u0001\u0000\u0000\u0000\u01b2\u01b3"+
		"\u0001\u0000\u0000\u0000\u01b3\u01b4\u0001\u0000\u0000\u0000\u01b4\u01b8"+
		"\u0005\u0001\u0000\u0000\u01b5\u01b7\u00030\u0018\u0000\u01b6\u01b5\u0001"+
		"\u0000\u0000\u0000\u01b7\u01ba\u0001\u0000\u0000\u0000\u01b8\u01b6\u0001"+
		"\u0000\u0000\u0000\u01b8\u01b9\u0001\u0000\u0000\u0000\u01b9\u01bb\u0001"+
		"\u0000\u0000\u0000\u01ba\u01b8\u0001\u0000\u0000\u0000\u01bb\u01bc\u0005"+
		"\u0002\u0000\u0000\u01bc\u01cf\u0001\u0000\u0000\u0000\u01bd\u01bf\u0005"+
		";\u0000\u0000\u01be\u01c0\u0003\u0012\t\u0000\u01bf\u01be\u0001\u0000"+
		"\u0000\u0000\u01bf\u01c0\u0001\u0000\u0000\u0000\u01c0\u01c1\u0001\u0000"+
		"\u0000\u0000\u01c1\u01c3\u0003:\u001d\u0000\u01c2\u01c4\u0003\u0018\f"+
		"\u0000\u01c3\u01c2\u0001\u0000\u0000\u0000\u01c3\u01c4\u0001\u0000\u0000"+
		"\u0000\u01c4\u01c5\u0001\u0000\u0000\u0000\u01c5\u01c9\u0005\u0001\u0000"+
		"\u0000\u01c6\u01c8\u00030\u0018\u0000\u01c7\u01c6\u0001\u0000\u0000\u0000"+
		"\u01c8\u01cb\u0001\u0000\u0000\u0000\u01c9\u01c7\u0001\u0000\u0000\u0000"+
		"\u01c9\u01ca\u0001\u0000\u0000\u0000\u01ca\u01cc\u0001\u0000\u0000\u0000"+
		"\u01cb\u01c9\u0001\u0000\u0000\u0000\u01cc\u01cd\u0005\u0002\u0000\u0000"+
		"\u01cd\u01cf\u0001\u0000\u0000\u0000\u01ce\u01aa\u0001\u0000\u0000\u0000"+
		"\u01ce\u01bd\u0001\u0000\u0000\u0000\u01cf/\u0001\u0000\u0000\u0000\u01d0"+
		"\u01d9\u0003\f\u0006\u0000\u01d1\u01d2\u0005=\u0000\u0000\u01d2\u01d3"+
		"\u0003\u008aE\u0000\u01d3\u01d4\u0005!\u0000\u0000\u01d4\u01d6\u0003:"+
		"\u001d\u0000\u01d5\u01d7\u0005\u0007\u0000\u0000\u01d6\u01d5\u0001\u0000"+
		"\u0000\u0000\u01d6\u01d7\u0001\u0000\u0000\u0000\u01d7\u01d9\u0001\u0000"+
		"\u0000\u0000\u01d8\u01d0\u0001\u0000\u0000\u0000\u01d8\u01d1\u0001\u0000"+
		"\u0000\u0000\u01d91\u0001\u0000\u0000\u0000\u01da\u01dc\u0005C\u0000\u0000"+
		"\u01db\u01da\u0001\u0000\u0000\u0000\u01db\u01dc\u0001\u0000\u0000\u0000"+
		"\u01dc\u01dd\u0001\u0000\u0000\u0000\u01dd\u01de\u0005<\u0000\u0000\u01de"+
		"\u01e0\u0003\u008aE\u0000\u01df\u01e1\u0003\u0012\t\u0000\u01e0\u01df"+
		"\u0001\u0000\u0000\u0000\u01e0\u01e1\u0001\u0000\u0000\u0000\u01e1\u01e2"+
		"\u0001\u0000\u0000\u0000\u01e2\u01e6\u0005\u0001\u0000\u0000\u01e3\u01e5"+
		"\u00034\u001a\u0000\u01e4\u01e3\u0001\u0000\u0000\u0000\u01e5\u01e8\u0001"+
		"\u0000\u0000\u0000\u01e6\u01e4\u0001\u0000\u0000\u0000\u01e6\u01e7\u0001"+
		"\u0000\u0000\u0000\u01e7\u01e9\u0001\u0000\u0000\u0000\u01e8\u01e6\u0001"+
		"\u0000\u0000\u0000\u01e9\u01ea\u0005\u0002\u0000\u0000\u01ea3\u0001\u0000"+
		"\u0000\u0000\u01eb\u01ed\u0005C\u0000\u0000\u01ec\u01eb\u0001\u0000\u0000"+
		"\u0000\u01ec\u01ed\u0001\u0000\u0000\u0000\u01ed\u01ee\u0001\u0000\u0000"+
		"\u0000\u01ee\u01ef\u0005A\u0000\u0000\u01ef\u01f0\u0003\u008aE\u0000\u01f0"+
		"\u01f1\u0005\b\u0000\u0000\u01f1\u01f4\u0003:\u001d\u0000\u01f2\u01f3"+
		"\u0005!\u0000\u0000\u01f3\u01f5\u0003h4\u0000\u01f4\u01f2\u0001\u0000"+
		"\u0000\u0000\u01f4\u01f5\u0001\u0000\u0000\u0000\u01f5\u01f7\u0001\u0000"+
		"\u0000\u0000\u01f6\u01f8\u0005\u0007\u0000\u0000\u01f7\u01f6\u0001\u0000"+
		"\u0000\u0000\u01f7\u01f8\u0001\u0000\u0000\u0000\u01f8\u01fb\u0001\u0000"+
		"\u0000\u0000\u01f9\u01fb\u0003\f\u0006\u0000\u01fa\u01ec\u0001\u0000\u0000"+
		"\u0000\u01fa\u01f9\u0001\u0000\u0000\u0000\u01fb5\u0001\u0000\u0000\u0000"+
		"\u01fc\u01fe\u0005C\u0000\u0000\u01fd\u01fc\u0001\u0000\u0000\u0000\u01fd"+
		"\u01fe\u0001\u0000\u0000\u0000\u01fe\u01ff\u0001\u0000\u0000\u0000\u01ff"+
		"\u0200\u0005=\u0000\u0000\u0200\u0202\u0003\u008aE\u0000\u0201\u0203\u0003"+
		"\u0012\t\u0000\u0202\u0201\u0001\u0000\u0000\u0000\u0202\u0203\u0001\u0000"+
		"\u0000\u0000\u0203\u0204\u0001\u0000\u0000\u0000\u0204\u0205\u0005!\u0000"+
		"\u0000\u0205\u0207\u0003:\u001d\u0000\u0206\u0208\u0005\u0007\u0000\u0000"+
		"\u0207\u0206\u0001\u0000\u0000\u0000\u0207\u0208\u0001\u0000\u0000\u0000"+
		"\u0208\u021c\u0001\u0000\u0000\u0000\u0209\u020b\u0005C\u0000\u0000\u020a"+
		"\u0209\u0001\u0000\u0000\u0000\u020a\u020b\u0001\u0000\u0000\u0000\u020b"+
		"\u020c\u0001\u0000\u0000\u0000\u020c\u020d\u0005=\u0000\u0000\u020d\u020f"+
		"\u0003\u008aE\u0000\u020e\u0210\u0003\u0012\t\u0000\u020f\u020e\u0001"+
		"\u0000\u0000\u0000\u020f\u0210\u0001\u0000\u0000\u0000\u0210\u0211\u0001"+
		"\u0000\u0000\u0000\u0211\u0212\u0005!\u0000\u0000\u0212\u0213\u0003:\u001d"+
		"\u0000\u0213\u0214\u0005>\u0000\u0000\u0214\u0215\u0005\u0013\u0000\u0000"+
		"\u0215\u0216\u0003\u008aE\u0000\u0216\u0217\u0005\u0013\u0000\u0000\u0217"+
		"\u0219\u0003h4\u0000\u0218\u021a\u0005\u0007\u0000\u0000\u0219\u0218\u0001"+
		"\u0000\u0000\u0000\u0219\u021a\u0001\u0000\u0000\u0000\u021a\u021c\u0001"+
		"\u0000\u0000\u0000\u021b\u01fd\u0001\u0000\u0000\u0000\u021b\u020a\u0001"+
		"\u0000\u0000\u0000\u021c7\u0001\u0000\u0000\u0000\u021d\u021f\u0005C\u0000"+
		"\u0000\u021e\u021d\u0001\u0000\u0000\u0000\u021e\u021f\u0001\u0000\u0000"+
		"\u0000\u021f\u0220\u0001\u0000\u0000\u0000\u0220\u0221\u0005?\u0000\u0000"+
		"\u0221\u0222\u0003\u008aE\u0000\u0222\u0223\u0005\b\u0000\u0000\u0223"+
		"\u0224\u0003:\u001d\u0000\u0224\u0225\u0005!\u0000\u0000\u0225\u0227\u0003"+
		"h4\u0000\u0226\u0228\u0005\u0007\u0000\u0000\u0227\u0226\u0001\u0000\u0000"+
		"\u0000\u0227\u0228\u0001\u0000\u0000\u0000\u02289\u0001\u0000\u0000\u0000"+
		"\u0229\u022a\u0006\u001d\uffff\uffff\u0000\u022a\u022f\u0003\u008aE\u0000"+
		"\u022b\u022c\u0005\u001d\u0000\u0000\u022c\u022d\u0003<\u001e\u0000\u022d"+
		"\u022e\u0005\u001e\u0000\u0000\u022e\u0230\u0001\u0000\u0000\u0000\u022f"+
		"\u022b\u0001\u0000\u0000\u0000\u022f\u0230\u0001\u0000\u0000\u0000\u0230"+
		"\u0245\u0001\u0000\u0000\u0000\u0231\u0232\u00057\u0000\u0000\u0232\u0234"+
		"\u0005\u0003\u0000\u0000\u0233\u0235\u0003<\u001e\u0000\u0234\u0233\u0001"+
		"\u0000\u0000\u0000\u0234\u0235\u0001\u0000\u0000\u0000\u0235\u0236\u0001"+
		"\u0000\u0000\u0000\u0236\u0237\u0005\u0004\u0000\u0000\u0237\u0238\u0005"+
		"\u000b\u0000\u0000\u0238\u0245\u0003:\u001d\u0007\u0239\u023a\u0005\u0003"+
		"\u0000\u0000\u023a\u023b\u0003<\u001e\u0000\u023b\u023c\u0005\u0004\u0000"+
		"\u0000\u023c\u0245\u0001\u0000\u0000\u0000\u023d\u023e\u0005\u0012\u0000"+
		"\u0000\u023e\u0245\u0003:\u001d\u0005\u023f\u0240\u0005\u0012\u0000\u0000"+
		"\u0240\u0241\u0005B\u0000\u0000\u0241\u0245\u0003:\u001d\u0004\u0242\u0243"+
		"\u0005;\u0000\u0000\u0243\u0245\u0003\u008aE\u0000\u0244\u0229\u0001\u0000"+
		"\u0000\u0000\u0244\u0231\u0001\u0000\u0000\u0000\u0244\u0239\u0001\u0000"+
		"\u0000\u0000\u0244\u023d\u0001\u0000\u0000\u0000\u0244\u023f\u0001\u0000"+
		"\u0000\u0000\u0244\u0242\u0001\u0000\u0000\u0000\u0245\u0252\u0001\u0000"+
		"\u0000\u0000\u0246\u0247\n\u0003\u0000\u0000\u0247\u0248\u0005\u0013\u0000"+
		"\u0000\u0248\u0251\u0003:\u001d\u0004\u0249\u024c\n\u0001\u0000\u0000"+
		"\u024a\u024b\u0005\f\u0000\u0000\u024b\u024d\u0003\u008aE\u0000\u024c"+
		"\u024a\u0001\u0000\u0000\u0000\u024d\u024e\u0001\u0000\u0000\u0000\u024e"+
		"\u024c\u0001\u0000\u0000\u0000\u024e\u024f\u0001\u0000\u0000\u0000\u024f"+
		"\u0251\u0001\u0000\u0000\u0000\u0250\u0246\u0001\u0000\u0000\u0000\u0250"+
		"\u0249\u0001\u0000\u0000\u0000\u0251\u0254\u0001\u0000\u0000\u0000\u0252"+
		"\u0250\u0001\u0000\u0000\u0000\u0252\u0253\u0001\u0000\u0000\u0000\u0253"+
		";\u0001\u0000\u0000\u0000\u0254\u0252\u0001\u0000\u0000\u0000\u0255\u025a"+
		"\u0003:\u001d\u0000\u0256\u0257\u0005\t\u0000\u0000\u0257\u0259\u0003"+
		":\u001d\u0000\u0258\u0256\u0001\u0000\u0000\u0000\u0259\u025c\u0001\u0000"+
		"\u0000\u0000\u025a\u0258\u0001\u0000\u0000\u0000\u025a\u025b\u0001\u0000"+
		"\u0000\u0000\u025b=\u0001\u0000\u0000\u0000\u025c\u025a\u0001\u0000\u0000"+
		"\u0000\u025d\u026c\u0003@ \u0000\u025e\u026c\u0003B!\u0000\u025f\u026c"+
		"\u0003D\"\u0000\u0260\u026c\u0003H$\u0000\u0261\u026c\u0003J%\u0000\u0262"+
		"\u026c\u0003L&\u0000\u0263\u026c\u0003N\'\u0000\u0264\u026c\u0003P(\u0000"+
		"\u0265\u026c\u0003R)\u0000\u0266\u026c\u0003T*\u0000\u0267\u026c\u0003"+
		"V+\u0000\u0268\u026c\u0003Z-\u0000\u0269\u026c\u0003\\.\u0000\u026a\u026c"+
		"\u0003^/\u0000\u026b\u025d\u0001\u0000\u0000\u0000\u026b\u025e\u0001\u0000"+
		"\u0000\u0000\u026b\u025f\u0001\u0000\u0000\u0000\u026b\u0260\u0001\u0000"+
		"\u0000\u0000\u026b\u0261\u0001\u0000\u0000\u0000\u026b\u0262\u0001\u0000"+
		"\u0000\u0000\u026b\u0263\u0001\u0000\u0000\u0000\u026b\u0264\u0001\u0000"+
		"\u0000\u0000\u026b\u0265\u0001\u0000\u0000\u0000\u026b\u0266\u0001\u0000"+
		"\u0000\u0000\u026b\u0267\u0001\u0000\u0000\u0000\u026b\u0268\u0001\u0000"+
		"\u0000\u0000\u026b\u0269\u0001\u0000\u0000\u0000\u026b\u026a\u0001\u0000"+
		"\u0000\u0000\u026c?\u0001\u0000\u0000\u0000\u026d\u026f\u0005@\u0000\u0000"+
		"\u026e\u0270\u0005B\u0000\u0000\u026f\u026e\u0001\u0000\u0000\u0000\u026f"+
		"\u0270\u0001\u0000\u0000\u0000\u0270\u0271\u0001\u0000\u0000\u0000\u0271"+
		"\u0274\u0003\u008aE\u0000\u0272\u0273\u0005\b\u0000\u0000\u0273\u0275"+
		"\u0003:\u001d\u0000\u0274\u0272\u0001\u0000\u0000\u0000\u0274\u0275\u0001"+
		"\u0000\u0000\u0000\u0275\u0276\u0001\u0000\u0000\u0000\u0276\u0277\u0005"+
		"!\u0000\u0000\u0277\u0279\u0003h4\u0000\u0278\u027a\u0005\u0007\u0000"+
		"\u0000\u0279\u0278\u0001\u0000\u0000\u0000\u0279\u027a\u0001\u0000\u0000"+
		"\u0000\u027aA\u0001\u0000\u0000\u0000\u027b\u027c\u0005A\u0000\u0000\u027c"+
		"\u027d\u0003\u008aE\u0000\u027d\u027e\u0005\b\u0000\u0000\u027e\u0281"+
		"\u0003:\u001d\u0000\u027f\u0280\u0005!\u0000\u0000\u0280\u0282\u0003h"+
		"4\u0000\u0281\u027f\u0001\u0000\u0000\u0000\u0281\u0282\u0001\u0000\u0000"+
		"\u0000\u0282\u0284\u0001\u0000\u0000\u0000\u0283\u0285\u0005\u0007\u0000"+
		"\u0000\u0284\u0283\u0001\u0000\u0000\u0000\u0284\u0285\u0001\u0000\u0000"+
		"\u0000\u0285C\u0001\u0000\u0000\u0000\u0286\u0287\u0003\u008aE\u0000\u0287"+
		"\u0288\u0003F#\u0000\u0288\u028a\u0003h4\u0000\u0289\u028b\u0005\u0007"+
		"\u0000\u0000\u028a\u0289\u0001\u0000\u0000\u0000\u028a\u028b\u0001\u0000"+
		"\u0000\u0000\u028bE\u0001\u0000\u0000\u0000\u028c\u028d\u0007\u0000\u0000"+
		"\u0000\u028dG\u0001\u0000\u0000\u0000\u028e\u028f\u0005U\u0000\u0000\u028f"+
		"\u0290\u0003\u008aE\u0000\u0290\u0291\u0003\u0084B\u0000\u0291I\u0001"+
		"\u0000\u0000\u0000\u0292\u0294\u0003h4\u0000\u0293\u0295\u0005\u0007\u0000"+
		"\u0000\u0294\u0293\u0001\u0000\u0000\u0000\u0294\u0295\u0001\u0000\u0000"+
		"\u0000\u0295K\u0001\u0000\u0000\u0000\u0296\u0298\u0005M\u0000\u0000\u0297"+
		"\u0299\u0003h4\u0000\u0298\u0297\u0001\u0000\u0000\u0000\u0298\u0299\u0001"+
		"\u0000\u0000\u0000\u0299\u029b\u0001\u0000\u0000\u0000\u029a\u029c\u0005"+
		"\u0007\u0000\u0000\u029b\u029a\u0001\u0000\u0000\u0000\u029b\u029c\u0001"+
		"\u0000\u0000\u0000\u029cM\u0001\u0000\u0000\u0000\u029d\u029e\u0005E\u0000"+
		"\u0000\u029e\u029f\u0003h4\u0000\u029f\u02a7\u0003\u0084B\u0000\u02a0"+
		"\u02a1\u0005F\u0000\u0000\u02a1\u02a2\u0005E\u0000\u0000\u02a2\u02a3\u0003"+
		"h4\u0000\u02a3\u02a4\u0003\u0084B\u0000\u02a4\u02a6\u0001\u0000\u0000"+
		"\u0000\u02a5\u02a0\u0001\u0000\u0000\u0000\u02a6\u02a9\u0001\u0000\u0000"+
		"\u0000\u02a7\u02a5\u0001\u0000\u0000\u0000\u02a7\u02a8\u0001\u0000\u0000"+
		"\u0000\u02a8\u02ac\u0001\u0000\u0000\u0000\u02a9\u02a7\u0001\u0000\u0000"+
		"\u0000\u02aa\u02ab\u0005F\u0000\u0000\u02ab\u02ad\u0003\u0084B\u0000\u02ac"+
		"\u02aa\u0001\u0000\u0000\u0000\u02ac\u02ad\u0001\u0000\u0000\u0000\u02ad"+
		"O\u0001\u0000\u0000\u0000\u02ae\u02af\u0003\u008aE\u0000\u02af\u02b0\u0005"+
		"\b\u0000\u0000\u02b0\u02b2\u0001\u0000\u0000\u0000\u02b1\u02ae\u0001\u0000"+
		"\u0000\u0000\u02b1\u02b2\u0001\u0000\u0000\u0000\u02b2\u02b3\u0001\u0000"+
		"\u0000\u0000\u02b3\u02b4\u0005G\u0000\u0000\u02b4\u02b5\u0003\u0084B\u0000"+
		"\u02b5Q\u0001\u0000\u0000\u0000\u02b6\u02b7\u0003\u008aE\u0000\u02b7\u02b8"+
		"\u0005\b\u0000\u0000\u02b8\u02ba\u0001\u0000\u0000\u0000\u02b9\u02b6\u0001"+
		"\u0000\u0000\u0000\u02b9\u02ba\u0001\u0000\u0000\u0000\u02ba\u02bb\u0001"+
		"\u0000\u0000\u0000\u02bb\u02bc\u0005H\u0000\u0000\u02bc\u02bd\u0003h4"+
		"\u0000\u02bd\u02be\u0003\u0084B\u0000\u02beS\u0001\u0000\u0000\u0000\u02bf"+
		"\u02c0\u0003\u008aE\u0000\u02c0\u02c1\u0005\b\u0000\u0000\u02c1\u02c3"+
		"\u0001\u0000\u0000\u0000\u02c2\u02bf\u0001\u0000\u0000\u0000\u02c2\u02c3"+
		"\u0001\u0000\u0000\u0000\u02c3\u02c4\u0001\u0000\u0000\u0000\u02c4\u02c5"+
		"\u0005I\u0000\u0000\u02c5\u02c6\u0003\u008aE\u0000\u02c6\u02c7\u0005J"+
		"\u0000\u0000\u02c7\u02c8\u0003h4\u0000\u02c8\u02c9\u0003\u0084B\u0000"+
		"\u02c9U\u0001\u0000\u0000\u0000\u02ca\u02cb\u0005N\u0000\u0000\u02cb\u02cc"+
		"\u0003h4\u0000\u02cc\u02d0\u0005\u0001\u0000\u0000\u02cd\u02cf\u0003X"+
		",\u0000\u02ce\u02cd\u0001\u0000\u0000\u0000\u02cf\u02d2\u0001\u0000\u0000"+
		"\u0000\u02d0\u02ce\u0001\u0000\u0000\u0000\u02d0\u02d1\u0001\u0000\u0000"+
		"\u0000\u02d1\u02d3\u0001\u0000\u0000\u0000\u02d2\u02d0\u0001\u0000\u0000"+
		"\u0000\u02d3\u02d4\u0005\u0002\u0000\u0000\u02d4W\u0001\u0000\u0000\u0000"+
		"\u02d5\u02d6\u0005O\u0000\u0000\u02d6\u02d7\u0003`0\u0000\u02d7\u02da"+
		"\u0007\u0001\u0000\u0000\u02d8\u02db\u0003>\u001f\u0000\u02d9\u02db\u0003"+
		"\u0084B\u0000\u02da\u02d8\u0001\u0000\u0000\u0000\u02da\u02d9\u0001\u0000"+
		"\u0000\u0000\u02dbY\u0001\u0000\u0000\u0000\u02dc\u02de\u0005K\u0000\u0000"+
		"\u02dd\u02df\u0003\u008aE\u0000\u02de\u02dd\u0001\u0000\u0000\u0000\u02de"+
		"\u02df\u0001\u0000\u0000\u0000\u02df\u02e1\u0001\u0000\u0000\u0000\u02e0"+
		"\u02e2\u0005\u0007\u0000\u0000\u02e1\u02e0\u0001\u0000\u0000\u0000\u02e1"+
		"\u02e2\u0001\u0000\u0000\u0000\u02e2[\u0001\u0000\u0000\u0000\u02e3\u02e5"+
		"\u0005L\u0000\u0000\u02e4\u02e6\u0003\u008aE\u0000\u02e5\u02e4\u0001\u0000"+
		"\u0000\u0000\u02e5\u02e6\u0001\u0000\u0000\u0000\u02e6\u02e8\u0001\u0000"+
		"\u0000\u0000\u02e7\u02e9\u0005\u0007\u0000\u0000\u02e8\u02e7\u0001\u0000"+
		"\u0000\u0000\u02e8\u02e9\u0001\u0000\u0000\u0000\u02e9]\u0001\u0000\u0000"+
		"\u0000\u02ea\u02eb\u0003\u0084B\u0000\u02eb_\u0001\u0000\u0000\u0000\u02ec"+
		"\u02ed\u00060\uffff\uffff\u0000\u02ed\u030c\u0003\u008aE\u0000\u02ee\u030c"+
		"\u00053\u0000\u0000\u02ef\u030c\u0003\u0088D\u0000\u02f0\u02f1\u0003\u008a"+
		"E\u0000\u02f1\u02f3\u0005\u0003\u0000\u0000\u02f2\u02f4\u0003b1\u0000"+
		"\u02f3\u02f2\u0001\u0000\u0000\u0000\u02f3\u02f4\u0001\u0000\u0000\u0000"+
		"\u02f4\u02f5\u0001\u0000\u0000\u0000\u02f5\u02f6\u0005\u0004\u0000\u0000"+
		"\u02f6\u030c\u0001\u0000\u0000\u0000\u02f7\u02f8\u0005@\u0000\u0000\u02f8"+
		"\u02fb\u0003\u008aE\u0000\u02f9\u02fa\u0005\b\u0000\u0000\u02fa\u02fc"+
		"\u0003:\u001d\u0000\u02fb\u02f9\u0001\u0000\u0000\u0000\u02fb\u02fc\u0001"+
		"\u0000\u0000\u0000\u02fc\u02ff\u0001\u0000\u0000\u0000\u02fd\u02fe\u0005"+
		"E\u0000\u0000\u02fe\u0300\u0003h4\u0000\u02ff\u02fd\u0001\u0000\u0000"+
		"\u0000\u02ff\u0300\u0001\u0000\u0000\u0000\u0300\u030c\u0001\u0000\u0000"+
		"\u0000\u0301\u0302\u0003\u008aE\u0000\u0302\u0303\u0005\u0001\u0000\u0000"+
		"\u0303\u0304\u0003f3\u0000\u0304\u0305\u0005\u0002\u0000\u0000\u0305\u030c"+
		"\u0001\u0000\u0000\u0000\u0306\u0308\u0005\u0003\u0000\u0000\u0307\u0309"+
		"\u0003b1\u0000\u0308\u0307\u0001\u0000\u0000\u0000\u0308\u0309\u0001\u0000"+
		"\u0000\u0000\u0309\u030a\u0001\u0000\u0000\u0000\u030a\u030c\u0005\u0004"+
		"\u0000\u0000\u030b\u02ec\u0001\u0000\u0000\u0000\u030b\u02ee\u0001\u0000"+
		"\u0000\u0000\u030b\u02ef\u0001\u0000\u0000\u0000\u030b\u02f0\u0001\u0000"+
		"\u0000\u0000\u030b\u02f7\u0001\u0000\u0000\u0000\u030b\u0301\u0001\u0000"+
		"\u0000\u0000\u030b\u0306\u0001\u0000\u0000\u0000\u030c\u0312\u0001\u0000"+
		"\u0000\u0000\u030d\u030e\n\u0005\u0000\u0000\u030e\u030f\u0005/\u0000"+
		"\u0000\u030f\u0311\u0003`0\u0006\u0310\u030d\u0001\u0000\u0000\u0000\u0311"+
		"\u0314\u0001\u0000\u0000\u0000\u0312\u0310\u0001\u0000\u0000\u0000\u0312"+
		"\u0313\u0001\u0000\u0000\u0000\u0313a\u0001\u0000\u0000\u0000\u0314\u0312"+
		"\u0001\u0000\u0000\u0000\u0315\u031a\u0003`0\u0000\u0316\u0317\u0005\t"+
		"\u0000\u0000\u0317\u0319\u0003`0\u0000\u0318\u0316\u0001\u0000\u0000\u0000"+
		"\u0319\u031c\u0001\u0000\u0000\u0000\u031a\u0318\u0001\u0000\u0000\u0000"+
		"\u031a\u031b\u0001\u0000\u0000\u0000\u031bc\u0001\u0000\u0000\u0000\u031c"+
		"\u031a\u0001\u0000\u0000\u0000\u031d\u031e\u0003\u008aE\u0000\u031e\u031f"+
		"\u0005\b\u0000\u0000\u031f\u0320\u0003`0\u0000\u0320\u0324\u0001\u0000"+
		"\u0000\u0000\u0321\u0324\u0003\u008aE\u0000\u0322\u0324\u0005/\u0000\u0000"+
		"\u0323\u031d\u0001\u0000\u0000\u0000\u0323\u0321\u0001\u0000\u0000\u0000"+
		"\u0323\u0322\u0001\u0000\u0000\u0000\u0324e\u0001\u0000\u0000\u0000\u0325"+
		"\u032a\u0003d2\u0000\u0326\u0327\u0005\t\u0000\u0000\u0327\u0329\u0003"+
		"d2\u0000\u0328\u0326\u0001\u0000\u0000\u0000\u0329\u032c\u0001\u0000\u0000"+
		"\u0000\u032a\u0328\u0001\u0000\u0000\u0000\u032a\u032b\u0001\u0000\u0000"+
		"\u0000\u032bg\u0001\u0000\u0000\u0000\u032c\u032a\u0001\u0000\u0000\u0000"+
		"\u032d\u032e\u00064\uffff\uffff\u0000\u032e\u0354\u0003l6\u0000\u032f"+
		"\u0330\u0007\u0002\u0000\u0000\u0330\u0354\u0003h4\u0015\u0331\u0332\u0005"+
		"Q\u0000\u0000\u0332\u0354\u0003h4\u0014\u0333\u0334\u0005R\u0000\u0000"+
		"\u0334\u0354\u0003h4\u0013\u0335\u0354\u0003n7\u0000\u0336\u0354\u0003"+
		"p8\u0000\u0337\u0338\u0005E\u0000\u0000\u0338\u0339\u0003h4\u0000\u0339"+
		"\u0341\u0003\u0084B\u0000\u033a\u033b\u0005F\u0000\u0000\u033b\u033c\u0005"+
		"E\u0000\u0000\u033c\u033d\u0003h4\u0000\u033d\u033e\u0003\u0084B\u0000"+
		"\u033e\u0340\u0001\u0000\u0000\u0000\u033f\u033a\u0001\u0000\u0000\u0000"+
		"\u0340\u0343\u0001\u0000\u0000\u0000\u0341\u033f\u0001\u0000\u0000\u0000"+
		"\u0341\u0342\u0001\u0000\u0000\u0000\u0342\u0346\u0001\u0000\u0000\u0000"+
		"\u0343\u0341\u0001\u0000\u0000\u0000\u0344\u0345\u0005F\u0000\u0000\u0345"+
		"\u0347\u0003\u0084B\u0000\u0346\u0344\u0001\u0000\u0000\u0000\u0346\u0347"+
		"\u0001\u0000\u0000\u0000\u0347\u0354\u0001\u0000\u0000\u0000\u0348\u0349"+
		"\u0005N\u0000\u0000\u0349\u034a\u0003h4\u0000\u034a\u034e\u0005\u0001"+
		"\u0000\u0000\u034b\u034d\u0003j5\u0000\u034c\u034b\u0001\u0000\u0000\u0000"+
		"\u034d\u0350\u0001\u0000\u0000\u0000\u034e\u034c\u0001\u0000\u0000\u0000"+
		"\u034e\u034f\u0001\u0000\u0000\u0000\u034f\u0351\u0001\u0000\u0000\u0000"+
		"\u0350\u034e\u0001\u0000\u0000\u0000\u0351\u0352\u0005\u0002\u0000\u0000"+
		"\u0352\u0354\u0001\u0000\u0000\u0000\u0353\u032d\u0001\u0000\u0000\u0000"+
		"\u0353\u032f\u0001\u0000\u0000\u0000\u0353\u0331\u0001\u0000\u0000\u0000"+
		"\u0353\u0333\u0001\u0000\u0000\u0000\u0353\u0335\u0001\u0000\u0000\u0000"+
		"\u0353\u0336\u0001\u0000\u0000\u0000\u0353\u0337\u0001\u0000\u0000\u0000"+
		"\u0353\u0348\u0001\u0000\u0000\u0000\u0354\u039f\u0001\u0000\u0000\u0000"+
		"\u0355\u0356\n\u0012\u0000\u0000\u0356\u0357\u0007\u0003\u0000\u0000\u0357"+
		"\u039e\u0003h4\u0013\u0358\u0359\n\u0011\u0000\u0000\u0359\u035a\u0007"+
		"\u0004\u0000\u0000\u035a\u039e\u0003h4\u0012\u035b\u035c\n\u0010\u0000"+
		"\u0000\u035c\u035d\u0007\u0005\u0000\u0000\u035d\u039e\u0003h4\u0011\u035e"+
		"\u035f\n\u000f\u0000\u0000\u035f\u0360\u0007\u0006\u0000\u0000\u0360\u039e"+
		"\u0003h4\u0010\u0361\u0362\n\u000e\u0000\u0000\u0362\u0363\u0007\u0007"+
		"\u0000\u0000\u0363\u039e\u0003h4\u000f\u0364\u0365\n\r\u0000\u0000\u0365"+
		"\u0366\u0005\u0012\u0000\u0000\u0366\u039e\u0003h4\u000e\u0367\u0368\n"+
		"\f\u0000\u0000\u0368\u0369\u0005\u0014\u0000\u0000\u0369\u039e\u0003h"+
		"4\r\u036a\u036b\n\u000b\u0000\u0000\u036b\u036c\u0005\u0013\u0000\u0000"+
		"\u036c\u039e\u0003h4\f\u036d\u036e\n\n\u0000\u0000\u036e\u036f\u0005\u0018"+
		"\u0000\u0000\u036f\u039e\u0003h4\u000b\u0370\u0371\n\t\u0000\u0000\u0371"+
		"\u0372\u0005\u0019\u0000\u0000\u0372\u039e\u0003h4\n\u0373\u0374\n\u0007"+
		"\u0000\u0000\u0374\u0375\u0005-\u0000\u0000\u0375\u039e\u0003h4\b\u0376"+
		"\u0377\n\u0006\u0000\u0000\u0377\u0378\u0005,\u0000\u0000\u0378\u0379"+
		"\u0003h4\u0000\u0379\u037a\u0005\b\u0000\u0000\u037a\u037b\u0003h4\u0007"+
		"\u037b\u039e\u0001\u0000\u0000\u0000\u037c\u037d\n\u0005\u0000\u0000\u037d"+
		"\u037e\u0007\b\u0000\u0000\u037e\u039e\u0003h4\u0006\u037f\u0380\n\u001a"+
		"\u0000\u0000\u0380\u0383\u0005\n\u0000\u0000\u0381\u0384\u0003\u008aE"+
		"\u0000\u0382\u0384\u0005Y\u0000\u0000\u0383\u0381\u0001\u0000\u0000\u0000"+
		"\u0383\u0382\u0001\u0000\u0000\u0000\u0384\u039e\u0001\u0000\u0000\u0000"+
		"\u0385\u0386\n\u0019\u0000\u0000\u0386\u0387\u0005\n\u0000\u0000\u0387"+
		"\u0388\u0003\u008aE\u0000\u0388\u038a\u0005\u0003\u0000\u0000\u0389\u038b"+
		"\u0003\u0086C\u0000\u038a\u0389\u0001\u0000\u0000\u0000\u038a\u038b\u0001"+
		"\u0000\u0000\u0000\u038b\u038c\u0001\u0000\u0000\u0000\u038c\u038d\u0005"+
		"\u0004\u0000\u0000\u038d\u039e\u0001\u0000\u0000\u0000\u038e\u038f\n\u0018"+
		"\u0000\u0000\u038f\u0390\u0005\u0005\u0000\u0000\u0390\u0391\u0003h4\u0000"+
		"\u0391\u0392\u0005\u0006\u0000\u0000\u0392\u039e\u0001\u0000\u0000\u0000"+
		"\u0393\u0394\n\u0017\u0000\u0000\u0394\u0396\u0005\u0003\u0000\u0000\u0395"+
		"\u0397\u0003\u0086C\u0000\u0396\u0395\u0001\u0000\u0000\u0000\u0396\u0397"+
		"\u0001\u0000\u0000\u0000\u0397\u0398\u0001\u0000\u0000\u0000\u0398\u039e"+
		"\u0005\u0004\u0000\u0000\u0399\u039a\n\u0016\u0000\u0000\u039a\u039e\u0007"+
		"\t\u0000\u0000\u039b\u039c\n\b\u0000\u0000\u039c\u039e\u0005,\u0000\u0000"+
		"\u039d\u0355\u0001\u0000\u0000\u0000\u039d\u0358\u0001\u0000\u0000\u0000"+
		"\u039d\u035b\u0001\u0000\u0000\u0000\u039d\u035e\u0001\u0000\u0000\u0000"+
		"\u039d\u0361\u0001\u0000\u0000\u0000\u039d\u0364\u0001\u0000\u0000\u0000"+
		"\u039d\u0367\u0001\u0000\u0000\u0000\u039d\u036a\u0001\u0000\u0000\u0000"+
		"\u039d\u036d\u0001\u0000\u0000\u0000\u039d\u0370\u0001\u0000\u0000\u0000"+
		"\u039d\u0373\u0001\u0000\u0000\u0000\u039d\u0376\u0001\u0000\u0000\u0000"+
		"\u039d\u037c\u0001\u0000\u0000\u0000\u039d\u037f\u0001\u0000\u0000\u0000"+
		"\u039d\u0385\u0001\u0000\u0000\u0000\u039d\u038e\u0001\u0000\u0000\u0000"+
		"\u039d\u0393\u0001\u0000\u0000\u0000\u039d\u0399\u0001\u0000\u0000\u0000"+
		"\u039d\u039b\u0001\u0000\u0000\u0000\u039e\u03a1\u0001\u0000\u0000\u0000"+
		"\u039f\u039d\u0001\u0000\u0000\u0000\u039f\u03a0\u0001\u0000\u0000\u0000"+
		"\u03a0i\u0001\u0000\u0000\u0000\u03a1\u039f\u0001\u0000\u0000\u0000\u03a2"+
		"\u03a3\u0005O\u0000\u0000\u03a3\u03a4\u0003`0\u0000\u03a4\u03a5\u0007"+
		"\u0001\u0000\u0000\u03a5\u03a6\u0003h4\u0000\u03a6k\u0001\u0000\u0000"+
		"\u0000\u03a7\u03b5\u0003\u008aE\u0000\u03a8\u03b5\u0005D\u0000\u0000\u03a9"+
		"\u03b5\u0003\u0088D\u0000\u03aa\u03ab\u0005\u0003\u0000\u0000\u03ab\u03ac"+
		"\u0003h4\u0000\u03ac\u03ad\u0005\u0004\u0000\u0000\u03ad\u03b5\u0001\u0000"+
		"\u0000\u0000\u03ae\u03b5\u0003\u0084B\u0000\u03af\u03b5\u0003r9\u0000"+
		"\u03b0\u03b5\u0003x<\u0000\u03b1\u03b5\u0003z=\u0000\u03b2\u03b5\u0003"+
		"~?\u0000\u03b3\u03b5\u0003\u0080@\u0000\u03b4\u03a7\u0001\u0000\u0000"+
		"\u0000\u03b4\u03a8\u0001\u0000\u0000\u0000\u03b4\u03a9\u0001\u0000\u0000"+
		"\u0000\u03b4\u03aa\u0001\u0000\u0000\u0000\u03b4\u03ae\u0001\u0000\u0000"+
		"\u0000\u03b4\u03af\u0001\u0000\u0000\u0000\u03b4\u03b0\u0001\u0000\u0000"+
		"\u0000\u03b4\u03b1\u0001\u0000\u0000\u0000\u03b4\u03b2\u0001\u0000\u0000"+
		"\u0000\u03b4\u03b3\u0001\u0000\u0000\u0000\u03b5m\u0001\u0000\u0000\u0000"+
		"\u03b6\u03b8\u0005Q\u0000\u0000\u03b7\u03b6\u0001\u0000\u0000\u0000\u03b7"+
		"\u03b8\u0001\u0000\u0000\u0000\u03b8\u03b9\u0001\u0000\u0000\u0000\u03b9"+
		"\u03bb\u0005\u0013\u0000\u0000\u03ba\u03bc\u0003\u000e\u0007\u0000\u03bb"+
		"\u03ba\u0001\u0000\u0000\u0000\u03bb\u03bc\u0001\u0000\u0000\u0000\u03bc"+
		"\u03bd\u0001\u0000\u0000\u0000\u03bd\u03c4\u0005\u0013\u0000\u0000\u03be"+
		"\u03bf\u0005\u000b\u0000\u0000\u03bf\u03c1\u0003:\u001d\u0000\u03c0\u03be"+
		"\u0001\u0000\u0000\u0000\u03c0\u03c1\u0001\u0000\u0000\u0000\u03c1\u03c2"+
		"\u0001\u0000\u0000\u0000\u03c2\u03c5\u0003\u0084B\u0000\u03c3\u03c5\u0003"+
		"h4\u0000\u03c4\u03c0\u0001\u0000\u0000\u0000\u03c4\u03c3\u0001\u0000\u0000"+
		"\u0000\u03c5o\u0001\u0000\u0000\u0000\u03c6\u03c7\u0005S\u0000\u0000\u03c7"+
		"\u03c8\u0003\u0084B\u0000\u03c8q\u0001\u0000\u0000\u0000\u03c9\u03ca\u0003"+
		"\u008aE\u0000\u03ca\u03cb\u0005\u0001\u0000\u0000\u03cb\u03cc\u0003t:"+
		"\u0000\u03cc\u03cd\u0005\u0002\u0000\u0000\u03cds\u0001\u0000\u0000\u0000"+
		"\u03ce\u03d3\u0003v;\u0000\u03cf\u03d0\u0005\t\u0000\u0000\u03d0\u03d2"+
		"\u0003v;\u0000\u03d1\u03cf\u0001\u0000\u0000\u0000\u03d2\u03d5\u0001\u0000"+
		"\u0000\u0000\u03d3\u03d1\u0001\u0000\u0000\u0000\u03d3\u03d4\u0001\u0000"+
		"\u0000\u0000\u03d4\u03d7\u0001\u0000\u0000\u0000\u03d5\u03d3\u0001\u0000"+
		"\u0000\u0000\u03d6\u03d8\u0005\t\u0000\u0000\u03d7\u03d6\u0001\u0000\u0000"+
		"\u0000\u03d7\u03d8\u0001\u0000\u0000\u0000\u03d8u\u0001\u0000\u0000\u0000"+
		"\u03d9\u03da\u0003\u008aE\u0000\u03da\u03db\u0005\b\u0000\u0000\u03db"+
		"\u03dc\u0003h4\u0000\u03dc\u03e1\u0001\u0000\u0000\u0000\u03dd\u03e1\u0003"+
		"\u008aE\u0000\u03de\u03df\u0005/\u0000\u0000\u03df\u03e1\u0003h4\u0000"+
		"\u03e0\u03d9\u0001\u0000\u0000\u0000\u03e0\u03dd\u0001\u0000\u0000\u0000"+
		"\u03e0\u03de\u0001\u0000\u0000\u0000\u03e1w\u0001\u0000\u0000\u0000\u03e2"+
		"\u03eb\u0005\u0005\u0000\u0000\u03e3\u03e8\u0003h4\u0000\u03e4\u03e5\u0005"+
		"\t\u0000\u0000\u03e5\u03e7\u0003h4\u0000\u03e6\u03e4\u0001\u0000\u0000"+
		"\u0000\u03e7\u03ea\u0001\u0000\u0000\u0000\u03e8\u03e6\u0001\u0000\u0000"+
		"\u0000\u03e8\u03e9\u0001\u0000\u0000\u0000\u03e9\u03ec\u0001\u0000\u0000"+
		"\u0000\u03ea\u03e8\u0001\u0000\u0000\u0000\u03eb\u03e3\u0001\u0000\u0000"+
		"\u0000\u03eb\u03ec\u0001\u0000\u0000\u0000\u03ec\u03ee\u0001\u0000\u0000"+
		"\u0000\u03ed\u03ef\u0005\t\u0000\u0000\u03ee\u03ed\u0001\u0000\u0000\u0000"+
		"\u03ee\u03ef\u0001\u0000\u0000\u0000\u03ef\u03f0\u0001\u0000\u0000\u0000"+
		"\u03f0\u03fe\u0005\u0006\u0000\u0000\u03f1\u03f2\u0005\u0005\u0000\u0000"+
		"\u03f2\u03f3\u0003h4\u0000\u03f3\u03f4\u0005I\u0000\u0000\u03f4\u03f5"+
		"\u0003\u008aE\u0000\u03f5\u03f6\u0005J\u0000\u0000\u03f6\u03f9\u0003h"+
		"4\u0000\u03f7\u03f8\u0005E\u0000\u0000\u03f8\u03fa\u0003h4\u0000\u03f9"+
		"\u03f7\u0001\u0000\u0000\u0000\u03f9\u03fa\u0001\u0000\u0000\u0000\u03fa"+
		"\u03fb\u0001\u0000\u0000\u0000\u03fb\u03fc\u0005\u0006\u0000\u0000\u03fc"+
		"\u03fe\u0001\u0000\u0000\u0000\u03fd\u03e2\u0001\u0000\u0000\u0000\u03fd"+
		"\u03f1\u0001\u0000\u0000\u0000\u03fey\u0001\u0000\u0000\u0000\u03ff\u0408"+
		"\u0005\u0001\u0000\u0000\u0400\u0405\u0003|>\u0000\u0401\u0402\u0005\t"+
		"\u0000\u0000\u0402\u0404\u0003|>\u0000\u0403\u0401\u0001\u0000\u0000\u0000"+
		"\u0404\u0407\u0001\u0000\u0000\u0000\u0405\u0403\u0001\u0000\u0000\u0000"+
		"\u0405\u0406\u0001\u0000\u0000\u0000\u0406\u0409\u0001\u0000\u0000\u0000"+
		"\u0407\u0405\u0001\u0000\u0000\u0000\u0408\u0400\u0001\u0000\u0000\u0000"+
		"\u0408\u0409\u0001\u0000\u0000\u0000\u0409\u040b\u0001\u0000\u0000\u0000"+
		"\u040a\u040c\u0005\t\u0000\u0000\u040b\u040a\u0001\u0000\u0000\u0000\u040b"+
		"\u040c\u0001\u0000\u0000\u0000\u040c\u040d\u0001\u0000\u0000\u0000\u040d"+
		"\u040e\u0005\u0002\u0000\u0000\u040e{\u0001\u0000\u0000\u0000\u040f\u0410"+
		"\u0003h4\u0000\u0410\u0411\u0005\b\u0000\u0000\u0411\u0412\u0003h4\u0000"+
		"\u0412}\u0001\u0000\u0000\u0000\u0413\u041c\u0005\u0001\u0000\u0000\u0414"+
		"\u0419\u0003h4\u0000\u0415\u0416\u0005\t\u0000\u0000\u0416\u0418\u0003"+
		"h4\u0000\u0417\u0415\u0001\u0000\u0000\u0000\u0418\u041b\u0001\u0000\u0000"+
		"\u0000\u0419\u0417\u0001\u0000\u0000\u0000\u0419\u041a\u0001\u0000\u0000"+
		"\u0000\u041a\u041d\u0001\u0000\u0000\u0000\u041b\u0419\u0001\u0000\u0000"+
		"\u0000\u041c\u0414\u0001\u0000\u0000\u0000\u041c\u041d\u0001\u0000\u0000"+
		"\u0000\u041d\u041f\u0001\u0000\u0000\u0000\u041e\u0420\u0005\t\u0000\u0000"+
		"\u041f\u041e\u0001\u0000\u0000\u0000\u041f\u0420\u0001\u0000\u0000\u0000"+
		"\u0420\u0421\u0001\u0000\u0000\u0000\u0421\u0422\u0005\u0002\u0000\u0000"+
		"\u0422\u007f\u0001\u0000\u0000\u0000\u0423\u042c\u0005\u0003\u0000\u0000"+
		"\u0424\u0429\u0003\u0082A\u0000\u0425\u0426\u0005\t\u0000\u0000\u0426"+
		"\u0428\u0003\u0082A\u0000\u0427\u0425\u0001\u0000\u0000\u0000\u0428\u042b"+
		"\u0001\u0000\u0000\u0000\u0429\u0427\u0001\u0000\u0000\u0000\u0429\u042a"+
		"\u0001\u0000\u0000\u0000\u042a\u042d\u0001\u0000\u0000\u0000\u042b\u0429"+
		"\u0001\u0000\u0000\u0000\u042c\u0424\u0001\u0000\u0000\u0000\u042c\u042d"+
		"\u0001\u0000\u0000\u0000\u042d\u042f\u0001\u0000\u0000\u0000\u042e\u0430"+
		"\u0005\t\u0000\u0000\u042f\u042e\u0001\u0000\u0000\u0000\u042f\u0430\u0001"+
		"\u0000\u0000\u0000\u0430\u0431\u0001\u0000\u0000\u0000\u0431\u0432\u0005"+
		"\u0004\u0000\u0000\u0432\u0081\u0001\u0000\u0000\u0000\u0433\u0434\u0003"+
		"\u008aE\u0000\u0434\u0435\u0005\b\u0000\u0000\u0435\u0436\u0003h4\u0000"+
		"\u0436\u0439\u0001\u0000\u0000\u0000\u0437\u0439\u0003h4\u0000\u0438\u0433"+
		"\u0001\u0000\u0000\u0000\u0438\u0437\u0001\u0000\u0000\u0000\u0439\u0083"+
		"\u0001\u0000\u0000\u0000\u043a\u043e\u0005\u0001\u0000\u0000\u043b\u043d"+
		"\u0003>\u001f\u0000\u043c\u043b\u0001\u0000\u0000\u0000\u043d\u0440\u0001"+
		"\u0000\u0000\u0000\u043e\u043c\u0001\u0000\u0000\u0000\u043e\u043f\u0001"+
		"\u0000\u0000\u0000\u043f\u0442\u0001\u0000\u0000\u0000\u0440\u043e\u0001"+
		"\u0000\u0000\u0000\u0441\u0443\u0003h4\u0000\u0442\u0441\u0001\u0000\u0000"+
		"\u0000\u0442\u0443\u0001\u0000\u0000\u0000\u0443\u0444\u0001\u0000\u0000"+
		"\u0000\u0444\u0445\u0005\u0002\u0000\u0000\u0445\u0085\u0001\u0000\u0000"+
		"\u0000\u0446\u044b\u0003h4\u0000\u0447\u0448\u0005\t\u0000\u0000\u0448"+
		"\u044a\u0003h4\u0000\u0449\u0447\u0001\u0000\u0000\u0000\u044a\u044d\u0001"+
		"\u0000\u0000\u0000\u044b\u0449\u0001\u0000\u0000\u0000\u044b\u044c\u0001"+
		"\u0000\u0000\u0000\u044c\u0087\u0001\u0000\u0000\u0000\u044d\u044b\u0001"+
		"\u0000\u0000\u0000\u044e\u044f\u0007\n\u0000\u0000\u044f\u0089\u0001\u0000"+
		"\u0000\u0000\u0450\u0451\u0005X\u0000\u0000\u0451\u008b\u0001\u0000\u0000"+
		"\u0000\u0090\u008f\u009c\u00a4\u00ad\u00b0\u00b8\u00ba\u00c1\u00c7\u00cd"+
		"\u00d1\u00d5\u00d8\u00dd\u00e1\u00e6\u00e9\u00f2\u0103\u010b\u0113\u011a"+
		"\u0123\u012b\u0130\u0136\u013c\u0143\u0146\u014b\u0153\u015d\u0164\u016a"+
		"\u016f\u0174\u017a\u0183\u0185\u018a\u018e\u0193\u0196\u019b\u019f\u01a4"+
		"\u01a8\u01ac\u01b2\u01b8\u01bf\u01c3\u01c9\u01ce\u01d6\u01d8\u01db\u01e0"+
		"\u01e6\u01ec\u01f4\u01f7\u01fa\u01fd\u0202\u0207\u020a\u020f\u0219\u021b"+
		"\u021e\u0227\u022f\u0234\u0244\u024e\u0250\u0252\u025a\u026b\u026f\u0274"+
		"\u0279\u0281\u0284\u028a\u0294\u0298\u029b\u02a7\u02ac\u02b1\u02b9\u02c2"+
		"\u02d0\u02da\u02de\u02e1\u02e5\u02e8\u02f3\u02fb\u02ff\u0308\u030b\u0312"+
		"\u031a\u0323\u032a\u0341\u0346\u034e\u0353\u0383\u038a\u0396\u039d\u039f"+
		"\u03b4\u03b7\u03bb\u03c0\u03c4\u03d3\u03d7\u03e0\u03e8\u03eb\u03ee\u03f9"+
		"\u03fd\u0405\u0408\u040b\u0419\u041c\u041f\u0429\u042c\u042f\u0438\u043e"+
		"\u0442\u044b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}