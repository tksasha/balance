class CashesController < ApplicationController
  #
  # TODO: try to use separate Resources for calculation and reloading tables with cashes and consolidates
  #
  # reload after change date or change Cash or Item
  #

  private
  def resource_params
    params.require(:cash).permit(:formula, :name)
  end
end
