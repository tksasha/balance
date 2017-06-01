class CashesController < ApplicationController
  include ActsAsRESTController

  private
  def collection
    @collection ||= Cash.order :name
  end

  def resource_params
    params.require(:cash).permit(:formula, :name)
  end
end
