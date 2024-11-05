# frozen_string_literal: true

RSpec.describe ItemsController do
  it { is_expected.to be_a(BaseController) }

  its(:default_url_options) { is_expected.to eq currency: 'uah' }

  describe '#scope', skip: 'private method' do
    before do
      allow(Item).to receive(:includes).with(:category) do
        double.tap do |a|
          allow(a).to receive(:order).with(date: :desc).and_return(:scope)
        end
      end
    end

    its(:scope) { is_expected.to eq :scope }
  end

  describe '#collection', skip: 'private method' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { is_expected.to eq :collection }
    end

    context do
      before do
        allow(subject).to receive_messages(params: :params, scope: :scope)

        allow(ItemSearcher).to receive(:search).with(:scope, :params).and_return(:collection)
      end

      its(:collection) { is_expected.to eq :collection }
    end
  end

  describe '#resource', skip: 'private method' do
    context 'when @resource is set' do
      before { subject.instance_variable_set(:@resource, :resource) }

      its(:resource) { is_expected.to eq :resource }
    end

    context 'when @resource is not set' do
      let(:params) { { id: 13 } }

      before do
        allow(subject).to receive(:params).and_return(params)

        allow(Item).to receive(:find).with(13).and_return(:resource)
      end

      its(:resource) { is_expected.to eq :resource }
    end
  end

  describe '#resource_params', skip: 'private method' do
    %w[uah usd eur].each do |currency|
      context "when currency is `#{ currency }`" do
        let(:params) do
          acp(
            currency:,
            item: { date: nil, formula: nil, category_id: nil, description: nil, tag_ids: [] }
          )
        end

        before { allow(subject).to receive(:params).and_return(params) }

        its(:resource_params) { is_expected.to eq params.require(:item).permit!.merge(currency:) }
      end
    end
  end

  describe '#build_resource', skip: 'private method' do
    before do
      allow(subject).to receive(:resource_params).and_return(:resource_params)

      allow(Item).to receive(:new).with(:resource_params).and_return(:resource)

      subject.send(:build_resource)
    end

    its(:resource) { is_expected.to eq :resource }
  end
end
