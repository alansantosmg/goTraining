using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace viewState
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {
            if(!Page.IsPostBack)
            {
                ViewState.Add("myViewStateT1", "");
            }
        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            string myValue = ViewState["myViewStateT1"].ToString();

            myValue += myTextBox.Text + " ";

            ViewState["myViewStateT1"] = myValue;

            resultLabel.Text = myValue;

            myTextBox.Text = "";

                    


        }
    }
}