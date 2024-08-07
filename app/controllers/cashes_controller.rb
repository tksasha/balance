# frozen_string_literal: true

class CashesController < Frontend::ApplicationController
  inherit_resources # TODO: !!! delme !!!

  actions :update

  respond_to :js

  def update
    update! do |_, failure|
      failure.js { render :edit }
    end
  end

  private

  def resource_params
    params.require(:cash).permit(:name, :formula)
  end

  def resource
    @resource ||= Cash.find(params[:id])
  end
end
