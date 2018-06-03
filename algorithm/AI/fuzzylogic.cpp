#include <algorithm>
#include "fuzzylogic.h"

/////////////////////////////////////////////////////////////////

//Initialise Fuzzy Rules

void initFuzzyRules(fuzzy_system_rec *fl) {
	
   const int
      // no_of_x_rules = 25,
      // no_of_theta_rules = 25;

      no_of_rules = 15;

   
   for (int i = 0;i < no_of_rules;i++) {
       fl->rules[i].inp_index[0] = in_yama_x;
       fl->rules[i].inp_index[1] = in_yama_y;
   }	

   fl->rules[0].inp_fuzzy_set[0] = in_nm;
   fl->rules[0].inp_fuzzy_set[1] = in_pm;
   fl->rules[0].out_fuzzy_set = out_ns;

   fl->rules[1].inp_fuzzy_set[0] = in_zr;
   fl->rules[1].inp_fuzzy_set[1] = in_pm;
   fl->rules[1].out_fuzzy_set = out_ps;

   fl->rules[2].inp_fuzzy_set[0] = in_pm;
   fl->rules[2].inp_fuzzy_set[1] = in_pm;
   fl->rules[2].out_fuzzy_set = out_pl;

   fl->rules[3].inp_fuzzy_set[0] = in_ns;
   fl->rules[3].inp_fuzzy_set[1] = in_ps;
   fl->rules[3].out_fuzzy_set = out_ns;

   fl->rules[4].inp_fuzzy_set[0] = in_ps;
   fl->rules[4].inp_fuzzy_set[1] = in_ps;
   fl->rules[4].out_fuzzy_set = out_pm;

   fl->rules[5].inp_fuzzy_set[0] = in_nm;
   fl->rules[5].inp_fuzzy_set[1] = in_zr;
   fl->rules[5].out_fuzzy_set = out_ns;

   fl->rules[6].inp_fuzzy_set[0] = in_zr;
   fl->rules[6].inp_fuzzy_set[1] = in_zr;
   fl->rules[6].out_fuzzy_set = out_zr;

   fl->rules[7].inp_fuzzy_set[0] = in_pm;
   fl->rules[7].inp_fuzzy_set[1] = in_zr;
   fl->rules[7].out_fuzzy_set = out_pm;

   fl->rules[8].inp_fuzzy_set[0] = in_ns;
   fl->rules[8].inp_fuzzy_set[1] = in_ns;
   fl->rules[8].out_fuzzy_set = out_nm;

   fl->rules[9].inp_fuzzy_set[0] = in_ps;
   fl->rules[9].inp_fuzzy_set[1] = in_ns;
   fl->rules[9].out_fuzzy_set = out_ps;

   fl->rules[10].inp_fuzzy_set[0] = in_nm;
   fl->rules[10].inp_fuzzy_set[1] = in_nm;
   fl->rules[10].out_fuzzy_set = out_nl;

   fl->rules[11].inp_fuzzy_set[0] = in_zr;
   fl->rules[11].inp_fuzzy_set[1] = in_nm;
   fl->rules[11].out_fuzzy_set = out_ns;


   fl->rules[12].inp_fuzzy_set[0] = in_pm;
   fl->rules[12].inp_fuzzy_set[1] = in_nm;
   fl->rules[12].out_fuzzy_set = out_ps;


   fl->rules[13].inp_fuzzy_set[0] = in_zr;
   fl->rules[13].inp_fuzzy_set[1] = in_ps;
   fl->rules[13].out_fuzzy_set = out_ps;


   fl->rules[14].inp_fuzzy_set[0] = in_zr;
   fl->rules[14].inp_fuzzy_set[1] = in_ns;
   fl->rules[14].out_fuzzy_set = out_ns;


   fl->rules[15].inp_fuzzy_set[0] = in_ns;
   fl->rules[15].inp_fuzzy_set[1] = in_zr;
   fl->rules[15].out_fuzzy_set = out_ns;

   fl->rules[16].inp_fuzzy_set[0] = in_ps;
   fl->rules[16].inp_fuzzy_set[1] = in_zr;
   fl->rules[16].out_fuzzy_set = out_ps;
//---------------------------------------------------------------------------- 	
//THETA vs. THETA_DOT	
//   
   // for (i = 0;i < no_of_theta_rules;i++) {
   //     fl->rules[i].inp_index[0] = in_theta;
   //     fl->rules[i].inp_index[1] = in_theta_dot;
   // }
      
   /* Regions for theta and theta_dot: */
   //sample only
   // fl->rules[0].inp_fuzzy_set[0] = in_nl;
   // fl->rules[0].inp_fuzzy_set[1] = in_nl;
   // fl->rules[0].out_fuzzy_set = out_nvl;
	
	
//----------------------------------------------------------------------------   
//X vs. X_DOT
//
 //   for (i = 0;i < no_of_x_rules;i++) {
 //   	  fl->rules[i + no_of_theta_rules].inp_index[0] = in_x;
 //   	  fl->rules[i + no_of_theta_rules].inp_index[1] = in_x_dot;
	// }
	  
	/* Regions for x and x_dot: */
   //sample only
   // fl->rules[25+0].inp_fuzzy_set[0] = in_nl;
   // fl->rules[25+0].inp_fuzzy_set[1] = in_nl;


   // fl->rules[25+24].out_fuzzy_set = out_nl;
      return;
}


void initMembershipFunctions(fuzzy_system_rec *fl) {
	
   /* The X membership functions */

   float x_base = ((A * THETA_MAX + B * THETA_DOT_MAX)-(A * THETA_MIN + B * THETA_DOT_MIN))/6;
  
   float y_base = ((C * X_MAX + D * X_DOT_MAX)-(C * X_MIN + D * X_DOT_MIN))/6;
  

   fl->inp_mem_fns[in_yama_x][in_nm] = init_trapz (x_base*(-3),x_base*(-2),x_base*(-2),x_base*(-1),regular_trapezoid);
   fl->inp_mem_fns[in_yama_x][in_ns] = init_trapz (x_base*(-2),x_base*(-1),x_base*(-1),x_base*(0),regular_trapezoid);
   fl->inp_mem_fns[in_yama_x][in_zr] = init_trapz (x_base*(-1),x_base*(0),x_base*(0),x_base*(1),regular_trapezoid);
   fl->inp_mem_fns[in_yama_x][in_ps] = init_trapz (x_base*(0),x_base*(1),x_base*(1),x_base*(2),regular_trapezoid);
   fl->inp_mem_fns[in_yama_x][in_pm] = init_trapz (x_base*(1),x_base*(2),x_base*(2),x_base*(3),regular_trapezoid);


   fl->inp_mem_fns[in_yama_y][in_nm] = init_trapz (y_base*(-3),y_base*(-2),y_base*(-2),y_base*(-1),regular_trapezoid);
   fl->inp_mem_fns[in_yama_y][in_ns] = init_trapz (y_base*(-2),y_base*(-1),y_base*(-1),y_base*(0),regular_trapezoid);
   fl->inp_mem_fns[in_yama_y][in_zr] = init_trapz (y_base*(-1),y_base*(0),y_base*(0),y_base*(1),regular_trapezoid);
   fl->inp_mem_fns[in_yama_y][in_ps] = init_trapz (y_base*(0),y_base*(1),y_base*(1),y_base*(2),regular_trapezoid);
   fl->inp_mem_fns[in_yama_y][in_pm] = init_trapz (y_base*(1),y_base*(2),y_base*(2),y_base*(3),regular_trapezoid);


   //Sample routines only, to give you an idea of what to do here
  	//~ fl->inp_mem_fns[in_x][in_neg] = init_trapz (-1.5,-0.5,0,0,left_trapezoid);
   //~ fl->inp_mem_fns[in_x][in_ze] = init_trapz (-1.5,-0.5,0.5,1.5,regular_trapezoid);
   //~ fl->inp_mem_fns[in_x][in_pos] = init_trapz (0.5,1.5,0,0,right_trapezoid);

   /* The X dot membership functions */
   //enter the appropriate membership function initialisations here 

   /* The theta membership functions */
   //enter the appropriate membership function initialisations here

   /* The theta dot membership functions */
   //enter the appropriate membership function initialisations here

 
   return;
}

void initFuzzySystem (fuzzy_system_rec *fl) {

   //Note: The settings of these parameters will depend upon your fuzzy system design
   // fl->no_of_inputs = 2;  /* Inputs are handled 2 at a time only */
   // fl->no_of_rules = 50;
   // fl->no_of_inp_regions = 5;
   // fl->no_of_outputs = 9;

   fl->no_of_inputs = 2;  /* Inputs are handled 2 at a time only */
   fl->no_of_rules = 17;
   fl->no_of_inp_regions = 5;
   fl->no_of_outputs = 7;
	
	
	//Sample only
	// fl->output_values [out_nvl]=-95.0;
	// fl->output_values [out_nl] = -85.0;


   fl->output_values [out_nl] = -40.0;
   fl->output_values [out_nm] = -30.0;
   fl->output_values [out_ns] = -10.0;
   fl->output_values [out_zr] = 0.0;
   fl->output_values [out_ps] = 10.0;
   fl->output_values [out_pm] = 30.0;
   fl->output_values [out_pl] = 40.0;
   

   fl->rules = (rule *) malloc ((size_t)(fl->no_of_rules*sizeof(rule)));
   initFuzzyRules(fl);
   initMembershipFunctions(fl);
   return;
}

//////////////////////////////////////////////////////////////////////////////

trapezoid init_trapz (float x1,float x2,float x3,float x4, trapz_type typ) {
	
   trapezoid trz;
   trz.a = x1;
   trz.b = x2;
   trz.c = x3;
   trz.d = x4;
   trz.tp = typ;
   switch (trz.tp) {
	   
      case regular_trapezoid:
         	 trz.l_slope = 1.0/(trz.b - trz.a);
         	 trz.r_slope = 1.0/(trz.c - trz.d);
         	 break;
	 
      case left_trapezoid:
      //add
         	 trz.r_slope = 1.0/(trz.a - trz.b);
         	 trz.l_slope = 0.0;
         	 break;
	 
      case right_trapezoid:
         	 trz.l_slope = 1.0/(trz.b - trz.a);
         	 trz.r_slope = 0.0;
         	 break;
   }  /* end switch  */
   
   return trz;
}  /* end function */

//////////////////////////////////////////////////////////////////////////////


//////////////////////////////////////////////////////////////////////////////
float trapz (float x, trapezoid trz) {
   switch (trz.tp) {
	   
      case left_trapezoid:
         	 if (x <= trz.a)
         	    return 1.0;
         	 if (x >= trz.b)
         	    return 0.0;
         	 /* a < x < b */
         	 return trz.r_slope * (x - trz.b);
	 
	 
      case right_trapezoid:
         	 if (x <= trz.a)
         	    return 0.0;
         	 if (x >= trz.b)
         	    return 1.0;
         	 /* a < x < b */
         	 return trz.l_slope * (x - trz.a);
	 
      case regular_trapezoid:
         	 if ((x <= trz.a) || (x >= trz.d))
         	    return 0.0;
         	 if ((x >= trz.b) && (x <= trz.c))
         	    return 1.0;
         	 if ((x >= trz.a) && (x <= trz.b))
         	    return trz.l_slope * (x - trz.a);
         	 if ((x >= trz.c) && (x <= trz.d))
         	    return  trz.r_slope * (x - trz.d);
         	    
	 }  /* End switch  */
	 
   return 0.0;  /* should not get to this point */
}  /* End function */

//////////////////////////////////////////////////////////////////////////////
float min_of(float values[],int no_of_inps) {
   int i;
   float val;
   val = values [0];
   for (i = 1;i < no_of_inps;i++) {
       if (values[i] < val)
	  val = values [i];
   }
   return val;
}



//////////////////////////////////////////////////////////////////////////////
float fuzzy_system (float inputs[],fuzzy_system_rec fz) {
   int i,j;
   short variable_index,fuzzy_set;
   float sum1 = 0.0,sum2 = 0.0,weight;
   float m_values[MAX_NO_OF_INPUTS];
	
   float yamakawaInputs[2];
   yamakawaInputs[0] = A*inputs[in_theta]+B*inputs[in_theta_dot];
   yamakawaInputs[1] = C*inputs[in_x]+D*inputs[in_x_dot]; 
   
   for (i = 0;i < fz.no_of_rules;i++) {
      for (j = 0;j < fz.no_of_inputs;j++) {
      	   variable_index = fz.rules[i].inp_index[j];
      	   fuzzy_set = fz.rules[i].inp_fuzzy_set[j];
      	   m_values[j] = trapz(yamakawaInputs[variable_index],fz.inp_mem_fns[variable_index][fuzzy_set]);
	   } /* end j  */
      
       weight = min_of (m_values,fz.no_of_inputs);
				
       sum1 += weight * fz.output_values[fz.rules[i].out_fuzzy_set];
       sum2 += weight;
   } /* end i  */
 
	
	if (fabs(sum2) < TOO_SMALL) {
      cout << "\r\nFLPRCS Error: Sum2 in fuzzy_system is 0.  Press key: " << endl;
      //~ getch();
      //~ exit(1);
      return 0.0;
   }
   return (sum1/sum2);
}  /* end fuzzy_system  */

//////////////////////////////////////////////////////////////////////////////
void free_fuzzy_rules (fuzzy_system_rec *fz) {
   if (fz->allocated){
	   free (fz->rules);
	}
	
   fz->allocated = false;
   return;
}

