using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace arrays1
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void submitButton1_Click(object sender, EventArgs e)
        {
            string[] myArray = new string[5] {"alan", "Santos", "teste", "cachimbo", "férias" };
            ViewState.Add("mystrings", myArray);


            resultLabel.Text = "Values added..."; 


        }

        protected void retrieveButton_Click(object sender, EventArgs e)
        {
            string[] myArray = (string[])ViewState["mystrings"];
            TextBox1.Text = myArray[0];
            TextBox2.Text = myArray[1];
            TextBox3.Text = myArray[2];
            TextBox4.Text = myArray[3];
            TextBox5.Text = myArray[4];
          
            

        }
    }
}