using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace simpleCalculator
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void addButton_Click(object sender, EventArgs e)
        {
            string firstValue = firstValueTextBox.Text;
            string secondValue = secondValueTextBox.Text;

            double result = double.Parse(firstValue) + double.Parse(secondValue);

            resultLabel.Text = result.ToString();


        }

        protected void subtractButton_Click(object sender, EventArgs e)
        {
            string firstValue = firstValueTextBox.Text;
            string secondValue = secondValueTextBox.Text;

            double result = double.Parse(firstValue) - double.Parse(secondValue);

            resultLabel.Text = result.ToString();
        }

        protected void multiplyButton_Click(object sender, EventArgs e)
        {
            string firstValue = firstValueTextBox.Text;
            string secondValue = secondValueTextBox.Text;

            double result = double.Parse(firstValue) * double.Parse(secondValue);

            resultLabel.Text = result.ToString();
        }

        protected void divideButton_Click(object sender, EventArgs e)
        {
            string firstValue = firstValueTextBox.Text;
            string secondValue = secondValueTextBox.Text;

            double result = double.Parse(firstValue) / double.Parse(secondValue);

            resultLabel.Text = result.ToString();
        }
    }
}