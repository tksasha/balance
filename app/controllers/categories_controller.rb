class CategoriesController < ApplicationController
  before_action :find_category, only: [:show, :edit, :destroy, :update]

  def index
    @categories = Category.all
  end

  def show
  end

  def new
    @category = Category.new
  end

  def create
    @category = Category.new category_params

    if @category.save
      redirect_to @category
    else
      render :new
    end
  end

  def destroy
    @category.destroy

    redirect_to :categories
  end

  def edit
  end

  def update
    if @category.update category_params
      redirect_to @category
    else
      render :edit
    end
  end

  private
  def find_category
    @category ||= Category.find params[:id]
  end

  def category_params
	  params.require(:category).permit(:name, :income, :slug, :visible)
  end
end
