# frozen_string_literal: true

RSpec.describe Backoffice::TagsController, type: :controller do
  describe '#category' do
    context do
      before { subject.instance_variable_set :@category, :category }

      its(:category) { should eq :category }
    end

    context do
      let(:params) { { category_id: 33 } }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(Category).to receive(:find).with(33).and_return(:category) }

      its(:category) { should eq :category }
    end

    its(:_helper_methods) { should include :category }
  end

  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { allow(subject).to receive(:category).and_return(:category) }

      before do
        allow(Backoffice::Tags::GetCollectionService)
          .to receive(:call)
          .with(:category)
          .and_return(:collection)
      end

      its(:collection) { should eq :collection }
    end
  end

  describe '#result' do
    context do
      before { subject.instance_variable_set :@result, :result }

      its(:result) { should eq :result }
    end

    context do
      before { allow(subject).to receive(:action_name).and_return(:action_name) }

      before { allow(subject).to receive(:category).and_return(:category) }

      before { allow(subject).to receive(:params).and_return(:params) }

      before do
        allow(Backoffice::Tags::GetResultService)
          .to receive(:call)
          .with(:action_name, :category, :params)
          .and_return(:result)
      end

      its(:result) { should eq :result }
    end
  end
end
