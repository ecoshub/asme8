// Generated from /Users/emre.ocak/go/src/test/antlr_asm/inst.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link instParser}.
 */
public interface instListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link instParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(instParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(instParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#instruction}.
	 * @param ctx the parse tree
	 */
	void enterInstruction(instParser.InstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#instruction}.
	 * @param ctx the parse tree
	 */
	void exitInstruction(instParser.InstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#line}.
	 * @param ctx the parse tree
	 */
	void enterLine(instParser.LineContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#line}.
	 * @param ctx the parse tree
	 */
	void exitLine(instParser.LineContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#inst_reg_reg}.
	 * @param ctx the parse tree
	 */
	void enterInst_reg_reg(instParser.Inst_reg_regContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#inst_reg_reg}.
	 * @param ctx the parse tree
	 */
	void exitInst_reg_reg(instParser.Inst_reg_regContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#inst_inm}.
	 * @param ctx the parse tree
	 */
	void enterInst_inm(instParser.Inst_inmContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#inst_inm}.
	 * @param ctx the parse tree
	 */
	void exitInst_inm(instParser.Inst_inmContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#inst_single_reg}.
	 * @param ctx the parse tree
	 */
	void enterInst_single_reg(instParser.Inst_single_regContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#inst_single_reg}.
	 * @param ctx the parse tree
	 */
	void exitInst_single_reg(instParser.Inst_single_regContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#inst_single_inm}.
	 * @param ctx the parse tree
	 */
	void enterInst_single_inm(instParser.Inst_single_inmContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#inst_single_inm}.
	 * @param ctx the parse tree
	 */
	void exitInst_single_inm(instParser.Inst_single_inmContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#symbol_multi}.
	 * @param ctx the parse tree
	 */
	void enterSymbol_multi(instParser.Symbol_multiContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#symbol_multi}.
	 * @param ctx the parse tree
	 */
	void exitSymbol_multi(instParser.Symbol_multiContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#symbol_single}.
	 * @param ctx the parse tree
	 */
	void enterSymbol_single(instParser.Symbol_singleContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#symbol_single}.
	 * @param ctx the parse tree
	 */
	void exitSymbol_single(instParser.Symbol_singleContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#reg}.
	 * @param ctx the parse tree
	 */
	void enterReg(instParser.RegContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#reg}.
	 * @param ctx the parse tree
	 */
	void exitReg(instParser.RegContext ctx);
	/**
	 * Enter a parse tree produced by {@link instParser#inm}.
	 * @param ctx the parse tree
	 */
	void enterInm(instParser.InmContext ctx);
	/**
	 * Exit a parse tree produced by {@link instParser#inm}.
	 * @param ctx the parse tree
	 */
	void exitInm(instParser.InmContext ctx);
}