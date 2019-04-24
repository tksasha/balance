class CashesController < ApplicationController
  before_action :set_variant, only: %i(index update)

  def create
    render :new, status: 422 unless resource.save
  end

  def update
    render :edit, status: 422 unless resource.update resource_params
  end

  def destroy
    resource.destroy
  end

  private
  def collection
    @collection ||= Cash.order :name
  end

  def resource_params
    params.require(:cash).permit(:formula, :name)
  end

  def set_variant
    request.variant = :report if params[:report].present?
  end

  def resource
    @resource ||= Cash.find params[:id]
  end

  def initialize_resource
    @resource = Cash.new
  end

  def build_resource
    @resource = Cash.new resource_params
  end
end
