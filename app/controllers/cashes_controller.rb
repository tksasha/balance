class CashesController < ApplicationController
  include ActsAsRESTController

  before_action :set_variant, only: [:index, :update]

  private
  def collection
    @collection ||= Cash.order :name
  end

  def resource_params
    params.require(:cash).permit(:formula, :name)
  end

  def set_variant
    request.variant = :report if params[:report]
  end
end
