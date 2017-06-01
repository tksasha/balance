class CashesController < ApplicationController
  #
  # TODO: try to use separate Resources for calculation and reloading tables with cashes and consolidates
  #
  # reload after change date or change Cash or Item
  #

  private
  def initialize_resource
    @cash = Cash.new
  end

  def build_resource
    @cash = Cash.new resource_params
  end

  def resource
    @cash ||= Cash.find params[:id]
  end

  def resource_params
    params.require(:cash).permit(:formula, :name)
  end

  def collection
    @collection ||= Cash.order :name
  end
end
