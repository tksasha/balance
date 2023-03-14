# frozen_string_literal: true

class CashesController < ApplicationController
  inherit_resources

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
