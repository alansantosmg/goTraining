using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace conditionalTernaryOperator
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            // resultLabel.Text = (TextBox1.Text == TextBox2.Text)
                

            resultLabel.Text = (myCheckBox.Checked)
                ? "yes, they are equal."
                : "no they are not equal.";
        }
    }
}