# frozen_string_literal: true

# rubocop:disable RSpec/HooksBeforeExamples
RSpec.shared_examples 'index' do |params|
  describe "GET #{ params[:uri] }" do
    include_examples 'format', params

    include_examples 'xhr?'

    before { get params[:uri], xhr: xhr? }

    it { should render_template :index }

    it_behaves_like 'media_type'
  end
end

RSpec.shared_examples 'new' do |params|
  describe "GET #{ params[:uri] }" do
    include_examples 'format', params

    include_examples 'xhr?'

    before { get params[:uri], xhr: xhr? }

    it { should render_template :new }

    it_behaves_like 'media_type'
  end
end

RSpec.shared_context 'format' do |params|
  let :format do
    #
    # '/items.json' -> 'json'
    #
    # '/items.js' -> 'js'
    #
    # '/items' -> 'html'
    #
    (md = /\.([[:alpha:]]+)$/.match(params[:uri])) ? md[1] : 'html'
  end
end

RSpec.shared_context 'xhr?' do
  let :xhr? do
    format == 'js'
  end
end

RSpec.shared_examples 'media_type' do
  context do
    subject { response }

    it { should have_http_status(:ok) }

    its :media_type do
      content_type = \
        case format
        when 'html'
          'text/html'
        when 'json'
          'application/json'
        when 'js'
          'text/javascript'
        end

      should eq content_type
    end
  end
end
# rubocop:enable RSpec/HooksBeforeExamples
